# マルチステージビルド-> build + nginx
FROM ubuntu:18.04 as build-stage
# 1. update
RUN apt-get update && apt-get  -y upgrade\
    && apt-get install -y curl
WORKDIR /root
# 2. create vue-emv
RUN curl -fsSL https://deb.nodesource.com/setup_16.x | bash -\
    && apt-get install -y nodejs\
    && curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add -\
    && echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list \
    && apt-get update && apt-get install -y yarn \
    && yarn global add @vue/cli@4.5.13
# 3. copy files
COPY . /root/front
WORKDIR /root/front
# 4. build vue-files
RUN yarn install \
    && yarn build


FROM nginx:1.21.0-alpine as production-stage
# 1. delete configs + make dir to put on contents
RUN rm -f /etc/nginx/conf.d/*.conf \
    && rm -f /etc/nginx/nginx.conf \
    && mkdir /home/www \
    && mkdir /home/www/contents
# 2. copy nginx.conf
COPY ./nginx/config/nginx.conf /etc/nginx/nginx.conf
# 3. move dist files which are built in build-stage
COPY --from=build-stage /root/front/dist /home/www/contents/dist
# 4. start nginx-server
CMD ["nginx", "-g", "daemon off;","-c","/etc/nginx/nginx.conf"]
