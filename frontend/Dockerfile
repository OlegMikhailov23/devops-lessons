# build
FROM node:16.20.0-alpine3.18 as builder
WORKDIR /usr/src/app
COPY package*.json ./
RUN npm ci
COPY . .
RUN npm run build
# runtime
FROM nginx:alpine

EXPOSE 80

COPY nginx.conf /etc/nginx/conf.d/default.conf

COPY --from=builder /usr/src/app/dist /usr/share/nginx/html
