/*
 * Strava API v3
 *
 * The [Swagger Playground](https://developers.strava.com/playground) is the easiest way to familiarize yourself with the Strava API by submitting HTTP requests and observing the responses before you write any client code. It will show what a response will look like with different endpoints depending on the authorization scope you receive from your athletes. To use the Playground, go to https://www.strava.com/settings/api and change your “Authorization Callback Domain” to developers.strava.com. Please note, we only support Swagger 2.0. There is a known issue where you can only select one scope at a time. For more information, please check the section “client code” at https://developers.strava.com/docs.
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ClubActivity struct {
	Athlete *MetaAthlete `json:"athlete,omitempty"`
	// The name of the activity
	Name string `json:"name,omitempty"`
	// The activity's distance, in meters
	Distance float32 `json:"distance,omitempty"`
	// The activity's moving time, in seconds
	MovingTime int32 `json:"moving_time,omitempty"`
	// The activity's elapsed time, in seconds
	ElapsedTime int32 `json:"elapsed_time,omitempty"`
	// The activity's total elevation gain.
	TotalElevationGain float32 `json:"total_elevation_gain,omitempty"`
	// Deprecated. Prefer to use sport_type
	Type_ *ActivityType `json:"type,omitempty"`
	SportType *SportType `json:"sport_type,omitempty"`
	// The activity's workout type
	WorkoutType int32 `json:"workout_type,omitempty"`
}
