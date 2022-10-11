FROM node:alpine
COPY . ./
RUN npm install
EXPOSE 8080
CMD ["npm", "start"]