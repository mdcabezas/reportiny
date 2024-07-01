
# ReporTiny


### GET /report
```bash
curl  -X  GET  \
  http://localhost/report
```
---
### POST /report
```bash
curl  -X  POST  \
  http://localhost/report \
  -H  "Content-Type: application/json"  \
  -d '{
    "image":  "https://placehold.co/200",
    "description":  "This is report",
    "user":  "mcabezas",
    "lat":  -35.290839,
    "lng":  -71.269964
  }'
```
---
### PUT /report/:id
```bash
curl -X PUT \
  http://localhost/report/cebaf48d-00d9-4d79-b915-8faa33a84e5a \
  -H "Content-Type: application/json" \
  -d '{
    "image":  "https://placehold.co/200",
    "description":  "This is report",
    "user":  "mcabezas",
    "lat":  -35.290839,
    "lng":  -71.269964
  }'
```
---
### DELETE /report/:id
```bash
curl -X DELETE \
  http://localhost/report/cebaf48d-00d9-4d79-b915-8faa33a84e5a
```
---