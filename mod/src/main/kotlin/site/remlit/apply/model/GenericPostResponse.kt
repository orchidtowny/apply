package site.remlit.apply.model

import kotlinx.serialization.Serializable

@Serializable
data class GenericPostResponse(
	val success: Boolean,
	val message: String
)
