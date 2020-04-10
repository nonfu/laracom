package laracom

import scala.concurrent.duration._

import io.gatling.core.Predef._
import io.gatling.http.Predef._
import io.gatling.jdbc.Predef._

class LoadTest extends Simulation {
	setUp(
        Scenarios.scn_SayHelloByUserId.inject(rampUsers(Conf.users) over (Scenarios.rampUpTimeSecs seconds)).protocols(Conf.httpConf)
    )
}
