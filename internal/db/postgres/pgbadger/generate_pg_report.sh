REPORTS_DIR="./internal/db/postgres/pgbadger/reports"
if [ ! -d "$REPORTS_DIR" ]; then
  mkdir -p "$REPORTS_DIR"
fi

echo '#!/bin/bash
LOG_DIR="/var/lib/postgresql/data/log"
REPORT_DIR="${LOG_DIR}/pgbadger_report"
if [ ! -d "$REPORT_DIR" ]; then
  mkdir "$REPORT_DIR"
fi
LATEST_LOG_FILE=$(ls -t $LOG_DIR/*.log | head -n 1)
if [ -z "$LATEST_LOG_FILE" ]; then
  echo "Не удалось найти лог-файл."
  exit 1
fi
# Get the current date and time for the report filename
DATE_TIME=$(date +"%Y-%m-%d_%H%M%S")
REPORT_FILE="${REPORT_DIR}/pgbadger_report_${DATE_TIME}.html"
pgbadger -f syslog "$LATEST_LOG_FILE" -o "$REPORT_FILE"
echo "Отчет сгенерирован: $REPORT_FILE"' | docker exec -i novamusic-postgres bash 
docker exec novamusic-postgres bash -c "ls -t /var/lib/postgresql/data/log/pgbadger_report/*.html | head -n 1" | xargs -I{} docker cp novamusic-postgres:{} ./internal/db/postgres/pgbadger/reports
