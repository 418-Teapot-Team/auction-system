# use node 16 alpine image as build image
FROM node:20-alpine
ADD .output /webapp/.output

EXPOSE 3000

WORKDIR /webapp/.output

ENV NUXT_HOST=0.0.0.0
ENV NUXT_PORT=3000

CMD ["node", "/webapp/.output/server/index.mjs"]