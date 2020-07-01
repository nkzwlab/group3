FROM node:8.16.0-alpine
WORKDIR /app

RUN apk add make curl
RUN npm update
RUN npx create-react-app .

CMD make start

RUN npm install -g react-scripts
RUN npx create-react-app /app
RUN make setup

CMD make start
