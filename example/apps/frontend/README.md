# Forge Example Frontend

This is a React application using MUI and TanStack Query to interact with the Forge-generated backend API.

## Development

### Prerequisites

- Node.js 18+
- npm or yarn

### Running locally

```bash
# Install dependencies
npm install

# Start development server
npm run dev
```

The app will be available at http://localhost:3000

### Building for production

```bash
npm run build
npm run preview
```

## Features

- **Tasks Management**: Create, read, update, and delete tasks
- **Projects Management**: Manage projects
- **MUI Components**: Modern Material Design UI
- **TanStack Query**: Efficient data fetching and caching

## Configuration

The app proxies API requests to the backend. In development, requests to `/api` are proxied to `http://localhost:8080`.

For production, configure your web server to proxy `/api` to your backend service.
