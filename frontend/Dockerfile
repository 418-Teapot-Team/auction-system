FROM node:lts-alpine

WORKDIR /app

COPY package*.json ./

RUN npm ci

COPY . .

RUN echo "BASE_API=http://142.93.172.206:6969/api/v1" > .env
RUN echo $BASE_API

RUN npm run build

# # customize port 80
# EXPOSE 80
# CMD ["sh", "-c", "NITRO_PORT=80 node .output/server/index.mjs"]

# default port 3000
EXPOSE 3000
CMD ["node", ".output/server/index.mjs"]