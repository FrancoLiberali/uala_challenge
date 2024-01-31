# uala_challenge

Coding challenge para Ualá

## Ejecución

### AWS

La solución del challenge se encuentra desplegada en AWS. A continuación se muestran los pasos para la ejecución de una demo utilizando curl:

Comienza a seguir a un usuario:

```bash
curl -X POST https://x7vywn4yub.execute-api.us-east-2.amazonaws.com/prod/user/1/follower/2
```

Crea un tweet:

```bash
curl -X POST https://x7vywn4yub.execute-api.us-east-2.amazonaws.com/prod/user/1/tweet -H 'Content-type: application/json' -d '{ "content": "este es un ejemplo" }'
```

Consulta el timeline:

```bash
curl https://x7vywn4yub.execute-api.us-east-2.amazonaws.com/prod/user/2/timeline
```

Para más información sobre los endpoints disponibles mira [api docs](#api-docs).

### Local

La misma demo puede ser ejecutada en la version local de la solución:

Lanza el servidor web:

```bash
docker compose -f "docker/docker-compose.yml" up -d
```

Comienza a seguir a un usuario:

```bash
curl -X POST localhost:8080/user/1/follower/2
```

Crea un tweet:

```bash
curl -X POST localhost:8080/user/1/tweet -H 'Content-type: application/json' -d '{ "content": "este es un ejemplo" }'
```

Consulta el timeline:

```bash
curl localhost:8080/user/2/timeline
```

## Api docs

La documentación de los endpoints disponibles se encuentra en el documento `swagger.json`. Puedes verlo de forma sencilla utilizando swagger-ui:

Corre el servidor web:

```bash
docker run -p 8081:8080 -e SWAGGER_JSON=/swagger.json -v ./swagger.json:/swagger.json swaggerapi/swagger-ui
```

Ingresa a [localhost:8081](localhost:8081) desde tu navegador.
