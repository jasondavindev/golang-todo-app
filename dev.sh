export PROJECT_NAME="golang_todo_app"
export PROJ_BASE=$(pwd)

function dkbuild {
  CD=$(pwd)
  cd $PROJ_BASE
  docker build -t $PROJECT_NAME -f docker/Dockerfile .
  exitcode=$?
  cd $CD
  return $exitcode
}

function dkup {
  CD=$(pwd)
  cd $PROJ_BASE
  PROJECT_NAME=$PROJECT_NAME docker-compose -f docker/docker-compose.yml up
  exitcode=$?
  cd $CD
  return $exitcode
}

function dkdown {
  CD=$(pwd)
  cd $PROJ_BASE
  docker-compose -f docker/docker-compose.yml down
  exitcode=$?
  cd $CD
  return $exitcode
}

function dk {
  docker exec -ti "$PROJECT_NAME"_app $@
}
