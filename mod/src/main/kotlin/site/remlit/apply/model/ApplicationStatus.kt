package site.remlit.apply.model

import kotlinx.serialization.Serializable

@Serializable
enum class ApplicationStatus {
	Pending,
	Rejected,
	Approved;

	fun fromInt(int: Int): ApplicationStatus =
		when (int) {
			1 -> ApplicationStatus.Rejected
			2 -> ApplicationStatus.Approved
			else -> ApplicationStatus.Pending
		}
}