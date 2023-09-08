variable "fqdn" {
  type    = string
  default = "troll.fejk.net"
}

variable "dcs" {
  type    = list(string)
  default = ["dc1", "devel"]
}


variable "image" {
  type = string
  default = "ghcr.io/theztd/troll:main"
}


job "__JOB_NAME__" {
  datacenters = var.dcs
  
  meta {
    fqdn = var.fqdn
    git = "github.com/theztd/troll"
    managed = "github-pipeline"
    image = var.image
  }

  group "fe" {
    count = 1

    network {
      mode = "bridge"

      dns {
        servers = ["172.17.0.1", "8.8.8.8", "1.1.1.1"]
      }
      
      port "app" { to = 8080 }
    }

    service {
      name = "${JOB}-app"
      provider = "nomad"

      tags = [
        "metrics=true",
        "metrics.path=/_healthz/metrics",
        "public",
        "traefik.enable=true",
        "traefik.http.routers.${NOMAD_JOB_NAME}-app.rule=Host(`${var.fqdn}`)"
      ]

      port = "app"

      check {
        name     = "${NOMAD_JOB_NAME} - alive"
        type     = "http"
        path     = "/v1/status"
        interval = "1m"
        timeout  = "10s"

        # Task should run 2m after deployment
        check_restart {
          limit           = 5
          grace           = "2m"
        }
      }

    }


    task "app" {

      driver = "docker"

      config {
        image      = var.image
        force_pull = true

        ports = ["app"]

        labels {
          group = "app"
        }
      }

      env {
        ADDRESS = ":8080"
        WAIT = 100
      }

      resources {
        cpu        = 100
        memory     = 64
        memory_max = 96
      }

    } # END task app


  } # END group FE

}
