
import scala.concurrent.duration._

import io.gatling.core.Predef._
import io.gatling.http.Predef._
import io.gatling.jdbc.Predef._

class SimulationForSimpleApp extends Simulation {

	val httpProtocol = http
		.baseUrl("http://localhost:8080")
		.inferHtmlResources()
		.acceptHeader("text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		.acceptEncodingHeader("gzip, deflate")
		.acceptLanguageHeader("en-US,en;q=0.9,hi;q=0.8")
		.doNotTrackHeader("1")
		.userAgentHeader("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36 Edg/92.0.902.84")

	val headers_0 = Map(
		"Cache-Control" -> "max-age=0",
		"Sec-Fetch-Dest" -> "document",
		"Sec-Fetch-Mode" -> "navigate",
		"Sec-Fetch-Site" -> "none",
		"Sec-Fetch-User" -> "?1",
		"Upgrade-Insecure-Requests" -> "1",
		"sec-ch-ua" -> """Chromium";v="92", " Not A;Brand";v="99", "Microsoft Edge";v="92""",
		"sec-ch-ua-mobile" -> "?0")

	val headers_1 = Map(
		"Accept" -> "image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8",
		"Cache-Control" -> "no-cache",
		"Pragma" -> "no-cache",
		"Sec-Fetch-Dest" -> "image",
		"Sec-Fetch-Mode" -> "no-cors",
		"Sec-Fetch-Site" -> "same-origin",
		"sec-ch-ua" -> """Chromium";v="92", " Not A;Brand";v="99", "Microsoft Edge";v="92""",
		"sec-ch-ua-mobile" -> "?0")

	val headers_2 = Map(
		"Sec-Fetch-Dest" -> "document",
		"Sec-Fetch-Mode" -> "navigate",
		"Sec-Fetch-Site" -> "none",
		"Sec-Fetch-User" -> "?1",
		"Upgrade-Insecure-Requests" -> "1",
		"sec-ch-ua" -> """Chromium";v="92", " Not A;Brand";v="99", "Microsoft Edge";v="92""",
		"sec-ch-ua-mobile" -> "?0")

	val headers_3 = Map(
		"Accept" -> "image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8",
		"Sec-Fetch-Dest" -> "image",
		"Sec-Fetch-Mode" -> "no-cors",
		"Sec-Fetch-Site" -> "same-origin",
		"sec-ch-ua" -> """Chromium";v="92", " Not A;Brand";v="99", "Microsoft Edge";v="92""",
		"sec-ch-ua-mobile" -> "?0")



	val scn = scenario("SimulationForSimpleApp")
		.exec(http("request_0")
			.get("/shiv")
			.headers(headers_0)
			.resources(http("request_1")
			.get("/favicon.ico")
			.headers(headers_1)))
		.pause(5)
		.exec(http("request_2")
			.get("/hi")
			.headers(headers_2)
			.resources(http("request_3")
			.get("/favicon.ico")
			.headers(headers_3)))
		.pause(18)
		.exec(http("request_4")
			.get("/s")
			.headers(headers_2)
			.resources(http("request_5")
			.get("/favicon.ico")
			.headers(headers_3)))
		.pause(6)
		.exec(http("request_6")
			.get("/h")
			.headers(headers_2)
			.resources(http("request_7")
			.get("/favicon.ico")
			.headers(headers_3)))
		.pause(3)
		.exec(http("request_8")
			.get("/hi")
			.headers(headers_2)
			.resources(http("request_9")
			.get("/favicon.ico")
			.headers(headers_3)))
		.pause(4)
		.exec(http("request_10")
			.get("/hi%20my")
			.headers(headers_2)
			.resources(http("request_11")
			.get("/favicon.ico")
			.headers(headers_3)))
		.pause(10)
		.exec(http("request_12")
			.get("/")
			.headers(headers_2)
			.resources(http("request_13")
			.get("/favicon.ico")
			.headers(headers_3)))
		.pause(5)
		.exec(http("request_14")
			.get("/check")
			.headers(headers_2)
			.resources(http("request_15")
			.get("/favicon.ico")
			.headers(headers_3)))

	setUp(scn.inject(atOnceUsers(1))).protocols(httpProtocol)
}