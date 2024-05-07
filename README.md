# GoServ

Basic implementation of a go-react-sqlite stack. Get started with this easy to use boilerplate.

Build react application with the following:

First modify the `.env` file in the client directory:

Append a `/api` at the end of the `REACT_APP_API_PREFIX`, then run:

`cd client && npm run build`

Then build the backend server with:

`go build`

You can host the static react contact with nginx, or any web server.
Serve the static context on the root directory, then setup a reverse proxy

An example nginx config for this:

```
server {
    listen 80;
    server_name <YOUR SERVER NAME/IP HERE>;

    location / {
        root <LOCATION OF REACT BUILD FILES>;
        try_files $uri /index.html;
    }

    location /api/ {
        proxy_pass http://localhost:8080;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

Run the backend server in a systemd file (or however you want to keep it alive forever). Then reload nginx config.

Good luck and **_happy building!_**