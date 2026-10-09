# Sayone Dey Portfolio

This portfolio uses Hugo for the static frontend and a Go standard-library server for the application layer.

## Run locally

```powershell
hugo
go run .\cmd\server
```

Then open `http://localhost:8080`. The Go service serves Hugo's `public` directory and provides:

- `GET /api/health` — service health check
- `GET /api/profile` — structured portfolio data
- `POST /api/contact` — validated contact form endpoint

The contact endpoint currently logs validated submissions. An email provider or database can be added later without changing the frontend contract.
