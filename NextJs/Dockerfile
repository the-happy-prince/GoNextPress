FROM node:current-alpine3.19
WORKDIR /app
COPY . .
RUN npm install
RUN npm run build
EXPOSE 3000
CMD npm run dev