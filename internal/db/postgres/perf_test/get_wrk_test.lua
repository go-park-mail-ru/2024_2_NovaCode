local wrk = require "wrk"

-- wrk.requests = 100000

wrk.method = "GET"
wrk.headers["Content-Type"] = "application/json"