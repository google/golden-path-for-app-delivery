#!/bin/bash -xe
cd frontend
skaffold render -p prod > ../prod.yaml
cd ..
conftest test -p policy ./prod.yaml
