## Back End for Judo Techniques

### How to run

- Clone the repo `git clone https://github.com/alexyak1/Techniques.git`
<details>
  <summary>Run with docker-compose</summary>

  ```bash
    docker-compose build
    docker-compose up
  ```
</details>

<details>
  <summary>Run with docker build and run</summary>

 - Build the Docker image `docker build -t techniques .` <br/>
 check that image created `docker images`. Should *techniques* exist
 - Run Container In the Background `docker run -p 8787:8787 -d --name techniques-container techniques` <br/>
 (Or run Container in interactive mode  `docker run -p 8787:8787 -it --name techniques-container techniques`) <br/>
 Port could be changed for any available ports. First port on local machine, second within container <br/>
 - Check that app is running. Open in browser localhost with selected port `http://localhost:8787/` <br/>

 - Kill Container `docker kill <container ID>` <br/> Check container ID `docker ps`
 - Remove container `docker rm -f techniques-container`
 </details>
<br/>

 ## Functionality
 if you run locally come to port <b>8787</b> `localhost:8787/`

 - Get all techniques: GET `localhost:8787/techniques`
 - Get all techniques with filter by belt: GET `localhost:8787/techniques?belt=yellow`
 - Get technique by id: GET `localhost:8787/technique{id}`
 - Add new technique: POST `localhost:8787/technique` with body example:
 ```json
 {
    "name": "Newly Created Post",
    "belt": "yellow",
    "image_url": "https://drive.google.com/url-path",
    "type": "Nage-Waza"
}
```
 - Delete technique: DELETE `localhost:8787/technique/{id}`
