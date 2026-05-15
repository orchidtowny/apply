package site.remlit.apply.model

import kotlinx.serialization.Serializable

@Serializable
data class ApprovedUsers(
	val usernames: List<String>
)
