package site.remlit.apply.model

import kotlinx.serialization.Serializable

@Serializable
data class Status(
	val username: String,
	val status: Int
)
