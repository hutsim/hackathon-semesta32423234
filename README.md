# 🕹️ Project Hackathon System Administrator SEMESTA Batch 5

Nama: Hary Miftah Fauzan

Asal Sekolah: SMKN 2 Tasikmalaya

## 📝 Deskripsi Project
Project ini dibangun 80 persen dibangun di lingkungan Google Cloud Provide (GCP), dengan topologi atau arsitektur sebagai berikut:

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
