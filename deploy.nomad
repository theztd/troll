variable "fqdn" {
  type = string
  default = "troll.fejk.net"
}

variable "dcs" {
  type = list(string)
  default = ["dc1", "devel"]
}


variable "image" {
  type = string
}


job "__JOB_NAME__" {
    datacenters = var.dcs

    group "fe" {
        count = 1

        network {
            port "app" { to = 8080 }
        }

        service {
            name = "${JOB}-http"

            tags = [
                "public",
                "traefik.enable=true",
                "traefik.http.routers.${NOMAD_JOB_NAME}-http.rule=Host(`${var.fqdn}`)",
                "traefik.http.routers.${NOMAD_JOB_NAME}-http.tls=true"
            ]

            port = "app"

            check {
                name = "${NOMAD_JOB_NAME} - alive"
                type = "http"
                path = "/v1/status"
                interval = "1m"
                timeout = "10s"

                # Task should run 2m after deployment
                check_restart {
                    limit = 5
                    grace = "2m"
                    ignore_warnings = true
                }
            }

         }

        task "app" {
          
          driver = "docker"

          config {
            image = var.image
            force_pull = true
            
            ports = ["app"]

            labels {
              group = "app"
            }
          }
          
          /*
          driver = "exec"

          config {
            command = "download"
            args = [
                "-name", "troll-1"
            ]
          }

          artifact {
            source = "https://gitlab.com/theztd/troll/-/package_files/47459529/download"
            options {
              checksum = "sha256:1b29427d34564b1ad14f7486577b0b94221dc07d2dfc92dbb942bd1d74d339ad"
            }
          }
          */

          env {
            ADDRESS = ":8080"
          }

          resources {
            cpu = 100
            memory = 64
            memory_max = 96
          }

        }

    } # END group FE

}