
https://github.com/jaegertracing/jaeger

https://www.jaegertracing.io/docs/2.13/getting-started/

https://github.com/jaegertracing/jaeger/blob/v2.13.0/examples/hotrod/README.md

start jaeger via docker
```text
docker run \
  --rm \
  --name jaeger \
  -p4318:4318 \
  -p16686:16686 \
  -p14268:14268 \
  cr.jaegertracing.io/jaegertracing/jaeger:2.13.0
  
  
  
docker run \
  --rm \
  --link jaeger \
  --env OTEL_EXPORTER_OTLP_ENDPOINT=http://jaeger:4318 \
  -p8080-8083:8080-8083 \
  jaegertracing/example-hotrod:latest \
  all
```