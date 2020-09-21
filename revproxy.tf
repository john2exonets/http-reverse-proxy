##
## revproxy
##
resource "docker_image" "revproxy" {
  name = "jdallen/revproxy:latest"
  keep_locally = true
}

resource "docker_container" "revproxy" {
  depends_on = [
    docker_container.my-demo-container,
    docker_network.demo-outside-net
  ]
  image = docker_image.revproxy.latest
  name = "revproxy"
  ports {
    protocol = "tcp"
    internal = 8080
    external = 8080
  }
  networks_advanced {
    name = "bridge"
  }
  networks_advanced {
    name = "demo-outside-net"
  }
  env = [
    "REVPROX_PORT=8080",
    "REVPROX_LOCAL_IP=0.0.0.0",
    "REVPROX_REMOTE=${local.demo_ips["demo-outside-net"]}:881"
  ]
  restart="always"
}

