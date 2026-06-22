FROM node:26-alpine

RUN npm install -g pnpm

WORKDIR /app

COPY package.json pnpm-workspace.yaml ./

RUN pnpm install

COPY . .

EXPOSE 8006

CMD ["pnpm", "exec", "tsx", "index.ts"]