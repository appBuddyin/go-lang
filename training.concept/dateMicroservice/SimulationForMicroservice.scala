
import scala.concurrent.duration._

import io.gatling.core.Predef._
import io.gatling.http.Predef._
import io.gatling.jdbc.Predef._

class SimulationForMicroservice extends Simulation {

	val httpProtocol = http
		.baseUrl("http://localhost:8080")
		.inferHtmlResources()
		.acceptHeader("text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		.acceptEncodingHeader("gzip, deflate")
		.acceptLanguageHeader("en-US,en;q=0.9,hi;q=0.8")
		.doNotTrackHeader("1")
		.upgradeInsecureRequestsHeader("1")
		.userAgentHeader("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36 Edg/92.0.902.84")

	val headers_0 = Map(
		"Cache-Control" -> "max-age=0",
		"Sec-Fetch-Dest" -> "document",
		"Sec-Fetch-Mode" -> "navigate",
		"Sec-Fetch-Site" -> "none",
		"Sec-Fetch-User" -> "?1",
		"sec-ch-ua" -> """Chromium";v="92", " Not A;Brand";v="99", "Microsoft Edge";v="92""",
		"sec-ch-ua-mobile" -> "?0")

	val headers_2 = Map(
		"Sec-Fetch-Dest" -> "document",
		"Sec-Fetch-Mode" -> "navigate",
		"Sec-Fetch-Site" -> "none",
		"Sec-Fetch-User" -> "?1",
		"sec-ch-ua" -> """Chromium";v="92", " Not A;Brand";v="99", "Microsoft Edge";v="92""",
		"sec-ch-ua-mobile" -> "?0")



	val scn = scenario("SimulationForMicroservice")
		.exec(http("request_0")
			.get("/status")
			.headers(headers_0))
		.pause(2)
		.exec(http("request_1")
			.get("/status")
			.headers(headers_0))
		.pause(17)
		.exec(http("request_2")
			.get("/get")
			.headers(headers_2))
		.pause(4)
		.exec(http("request_3")
			.get("/get")
			.headers(headers_0))
		.pause(122)
		.exec(http("request_4")
			.get("/get")
			.headers(headers_0)
			.resources(http("request_5")
			.get("/get")
			.headers(headers_0)))
		.pause(4)
		.exec(http("request_6")
			.get("/status")
			.headers(headers_2))
		.pause(4)
		.exec(http("request_7")
			.get("/get")
			.headers(headers_2))

	setUp(scn.inject(atOnceUsers(1))).protocols(httpProtocol)
}