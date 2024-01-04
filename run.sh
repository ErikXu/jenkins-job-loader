#!/bin/bash

./jenkins-job-loader load -d "${DOMAIN}" -u "${USER}" -p "${PASSWORD}" -f "${FOLDER}" -c "${CREDENTIALS_ID}"