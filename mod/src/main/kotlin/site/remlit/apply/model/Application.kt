package site.remlit.apply.model

import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

@Serializable
data class Application(
	val id: String,
	val username: String,
	val age: Int,
	@SerialName("where_did_you_find_the_server")
	val whereDidYouFindTheServer: String,
	val bio: String,
	val status: Int
)