FROM alpine:3.15

RUN apk add --no-cache tzdata

ARG CI_PROJECT_TITLE
ENV APP_NAME=${CI_PROJECT_TITLE}

WORKDIR /app
RUN echo -e "#!/bin/sh\n ./${APP_NAME} \$@" > ./entrypoint.sh && chmod +x ./entrypoint.sh
COPY build/${CI_PROJECT_TITLE} .

ENTRYPOINT ["./entrypoint.sh"]
