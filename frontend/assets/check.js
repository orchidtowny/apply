window.onload = () => {
    const form = document.getElementById("check-form")
    if (!(form instanceof HTMLFormElement)) {
        console.log("error with check-form, fail at !instanceof")
        return
    }

    function tryUsername(username) {
        window.fetch(
            "/api/status/"+username,
            { method: "GET" }
        ).then(async (response) => {
            console.log(response)

            if (response.status === 200) {
                const json = await response.json()

                let message = "Unknown"
                switch (json?.status) {
                    case 1:
                        message = "Application rejected."
                        break
                    case 2:
                        message = "Application approved."
                        break
                    default:
                        message = "Application still awaiting approval."
                        break
                }

                document.getElementById("checked-user").innerHTML = `
                    <p>${message}</p>
                `
            } else if (response.status === 404) {
                document.getElementById("checked-user").innerHTML = `
                    <p>Application not found.</p>
                `
            } else {
                document.getElementById("checked-user").innerHTML = `
                    <p>Something went wrong, try again later.</p>
                `
            }
        })
    }

    const query = new URLSearchParams(window.location.search)
    const existingUsername = query.get("username")

    if (existingUsername !== "" && existingUsername !== undefined) {
        tryUsername(existingUsername)
    }

    form.onsubmit = (ev) => {
        ev.preventDefault()

        const data = new FormData(event.target);
        const username = data.get("username")

        tryUsername(username)
    }
}