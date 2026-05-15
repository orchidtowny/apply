package site.remlit.apply.model

import kotlinx.serialization.Serializable

@Serializable
data class Info(
	val open: Boolean,
	val rules: Map<String, String>
)