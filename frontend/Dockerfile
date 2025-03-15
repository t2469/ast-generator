FROM node:22-alpine AS builder
WORKDIR /home/node/app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

FROM node:22-alpine
WORKDIR /home/node/app
COPY --from=builder /home/node/app/dist ./dist
RUN npm install -g serve
EXPOSE 5173
CMD ["serve", "-s", "dist", "-l", "5173"]
