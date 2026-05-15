package site.remlit.apply

import net.minecraftforge.fml.common.Mod
import net.minecraftforge.fml.event.lifecycle.FMLDedicatedServerSetupEvent
import thedarkcolour.kotlinforforge.forge.MOD_BUS
import thedarkcolour.kotlinforforge.forge.runForDist
import java.util.logging.LogManager
import java.util.logging.Logger

@Mod(Apply.ID)
object Apply {
	const val ID: String = "apply"

	val LOGGER: Logger = LogManager.getLogManager().getLogger(ID)

	init {
		LOGGER.info("Apply started!")

		val obj = runForDist(
			clientTarget = {},
			serverTarget = {
				MOD_BUS.addListener(::onServerSetup)
				"test"
			}
		)

		println(obj)
	}

	private fun onServerSetup(event: FMLDedicatedServerSetupEvent) {
		LOGGER.info { "Server setup hook hit" }
	}
}