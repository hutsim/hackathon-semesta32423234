# 🕹️ Project Hackathon System Administrator SEMESTA Batch 5

Nama: Hary Miftah Fauzan

Asal Sekolah: SMKN 2 Tasikmalaya

## 📝 Deskripsi Project
Project ini dibangun 80 persen di lingkungan Google Cloud Provide (GCP), dengan topologi atau arsitektur sebagai berikut:

![Teks Alt](topo-hackathon.png)

Terdapat 2 tipe:
1. User
    - User adalah orang yang mengakses website dari internet. Alur topologi bagi user yaitu:
        * User mengetikkan url https://semesta.unvizy.xyz
        * Kemudian oleh DNS akan diarahkan ke Ingress global https
        * Setelah diarahkan ke service app1
        * kemudian dari service app1 akan diarahkan ke dalam pod app1
        * jika user mengakses menggunakan endpoint /aboutus, maka pod app1 akan melakukan request ke service app2
        * oleh service app2 akan diarahkan ke pod app2

2. Developer 
    - adalah orang yang melakukan develope website tersebut. Alur untuk developer sebagai berikut:
        * Developer melakukan push source code ke github repository
        * Jenkins melakukan pull source code
        * Jenkins melakukan job pipeline, diantaranya:
            * checkout
            * test source code
            * build docker image
            * push docker image ke docker registry
            * deploy aplikasi ke kubernetes cluster

Server Jenkins diinstall menggunakan Google Cloud Compute, proses provisioning server menggunakan Terraform dan untuk setup server nya menggunakan Ansible

### Berikut tampilan hasil akses website nya:

- Akses https://semesta.unvizy.xyz
    ![Teks Alt](root.png)
- Akses https://semesta.unvizy.xyz/aboutus
    ![Teks Alt](aboutus.png)


## 🚀 Cara Menggunakan Project
### Kubernetes Cluster
Buat terlebih dahulu kubernetes cluster nya.
Saya di sini menggunakan Google Kubernetes Engine (GKE)
        
![Teks Alt](k8scluster.png)
        
### Persipkan server Jenkins. 
Server Jenkins dibangun menggunakan terraform, dan ansible

### Jenkins pipeline
![Teks Alt](pipeline.png)

## 🌐 Deskripsi Project Cisco Packet Tracer
Architectfure network dengan HA dan juga etherchannel. Tidak lupa access control list untuk melakukan filtering traffic.
Filtering traffic. digunakan agar PC2 tidak dapat mengakses website.
![Teks Alt](aksesPC2keWeb.png)
untuk ping, PC2 diizinkan melakukan ping ke server
web.
![Teks Alt](pingPC2AslikeServer.png)

Akses PC1 ke Website:
![Teks Alt](aksesPC1keWeb.png)

HA pada Router1
![Teks Alt](hsrp-router1.png)

HA pada Router2
![Teks Alt](hsrp-router2.png)