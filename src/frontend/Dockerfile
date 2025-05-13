FROM node:18

WORKDIR /app

COPY package*.json ./
RUN npm ci --ignore-scripts

COPY . .

RUN npm run dev
EXPOSE 3000
CMD ["npm", "start"]