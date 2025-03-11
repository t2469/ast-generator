FROM node:22-alpine AS builder

WORKDIR /home/node/app

COPY package*.json ./
RUN npm install

COPY . .

FROM node:22-alpine

WORKDIR /home/node/app

COPY --from=builder /home/node/app ./

RUN npm prune --production

EXPOSE 5173

CMD ["npm", "start"]
