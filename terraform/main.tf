provider "google" {
    credentials = file("/home/unvizy/.credentials/semesta.json")
    project = "semesta-390801"
    region = "asia-southeast1"
    zone = "asia-southeast1-a"
  
}

resource "google_compute_address" "vm_static_ip" {
    name = "semesta-static-ip"
  
}

resource "google_compute_instance" "vm_instance" {
    name = "semesta-instance"
    machine_type = "e2-medium"
    provisioner "local-exec" {
      command = "echo ${google_compute_instance.vm_instance.name}:  ${google_compute_instance.vm_instance.network_interface[0].access_config[0].nat_ip} >> ip_address.txt"
    }
    tags = ["http-server","https-server"]


    metadata = {
      ssh-keys = "unvizy:${file("~/.ssh/id_rsa.pub")}"
    }

    boot_disk {
        initialize_params {
          image = "ubuntu-os-cloud/ubuntu-2204-lts"
        }
      
    }
    network_interface {
      network = "default"
      access_config {
         nat_ip = google_compute_address.vm_static_ip.address        
      }
    }  
}

