FROM node:16

WORKDIR /usr/src/app
COPY package*.json ./
RUN npm install
ENV LANDING_HOSTS_PATH=hosts.json

COPY . .

EXPOSE  8080

CMD [ "npm", "start" ]