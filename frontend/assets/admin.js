window.onload = () => {



    function changeStatus(id) {

    }

    function createEntry(data) {
        `
        <div class="adminApplicationEntry">
            <div>
                <b>Username</b>
                <span>${data.user}</span>
            </div>
            
            <div>
                <b>Age</b>
                <span>${data.age}</span>
            </div>
            
            <div>
                <b>Where did you find Orchid?</b>
                <span>${data.where_did_you_find_the_server}</span>
            </div>
            
            <div>
                <b>Tell us a little about yourself</b>
                <span>${data.bio}</span>
            </div>
        
            <label>
                Status
                <select id="status-select-${data.id}" onchange="changeStatus(${data.id})">
                    <option value="0" name="Pending" ${data.status === 0 ? "selected" : ""} />
                    <option value="1" name="Rejected" ${data.status === 1 ? "selected" : ""} />
                    <option value="2" name="Approved" ${data.status === 2 ? "selected" : ""} />
                </select>
            </label>
        </div>
        `
    }
}