### STAGE 1: BUILD ###
FROM node:15-alpine AS builder
WORKDIR /usr/src/app
COPY . .
RUN npm install
RUN npm run build


### STAGE 2: RUN ###
FROM nginx:1.19.7-alpine
COPY nginx.conf /etc/nginx/nginx.conf
COPY --from=builder /usr/src/app/dist/web /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]


