# Stage 1: Build
FROM node:20-alpine AS builder

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm ci

COPY . .
RUN npm run build

# Stage 2: Production runner
FROM node:20-alpine AS runner

WORKDIR /app

# Copy the compiled Node.js app (adapter-node bundles most things)
COPY --from=builder /app/build ./build

EXPOSE 3000

ENV PORT=3000
ENV NODE_ENV=production
ENV HOST=0.0.0.0

CMD ["node", "build/index.js"]