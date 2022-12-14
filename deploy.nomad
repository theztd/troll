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
}


job "__JOB_NAME__" {
  datacenters = var.dcs

  group "fe" {
    count = 1

    network {
      mode = "bridge"

      dns {
        servers = ["172.17.0.1", "8.8.8.8", "1.1.1.1"]
      }
      
      port "app" { to = 8080 }
      port "http" { to = 80 }
    }

    service {
      name = "${JOB}-http"

      tags = [
        "public",
        "traefik.enable=true",
        "traefik.http.routers.${NOMAD_JOB_NAME}-http.rule=Host(`http-${var.fqdn}`)"
        //"traefik.http.routers.${NOMAD_JOB_NAME}-http.tls=true"
      ]

      port = "http"
    }

    service {
      name = "${JOB}-app"

      tags = [
        "public",
        "traefik.enable=true",
        "traefik.http.routers.${NOMAD_JOB_NAME}-app.rule=Host(`${var.fqdn}`)"
        //"traefik.http.routers.${NOMAD_JOB_NAME}-http.tls=true"
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
          ignore_warnings = true
        }
      }

    }

    task "nginx" {
      driver = "docker"

      config {
        image = "nginx:1.21"

        volumes = [
          "local:/etc/nginx/conf.d",
        ]

        ports = ["http"]
      }

      template {
        destination = "local/default.conf"
        perms       = "644"
        data        = file("nginx.conf")
      }

      # Resources:    https://www.nomadproject.io/docs/job-specification/resources
      resources {
        cpu        = 100 # MHz
        memory     = 16  # MB
        memory_max = 64  #MB
      }


      kill_timeout = "10s"
    }
    # END NGinx task





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
