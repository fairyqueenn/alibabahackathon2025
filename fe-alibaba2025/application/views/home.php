<!DOCTYPE html>
<html lang="en">

<head>
    <?php $this->load->view('template/head/meta', ['title' => 'Home', 'url' => site_url()], FALSE); ?>
    <?php $this->load->view('template/head/mandatory_style', NULL, FALSE); ?>
    <style>
        .card-product *:hover {
            cursor: pointer !important;
        }
    </style>
</head>

<body>
    <div class="wrapper">
        <div class="wrapper-header">
            <div class="d-flex justify-content-between">
                <div>
                    <img src="<?= base_url("assets/img/logo.png"); ?>" alt="" style="width: 25px; height: 25px;">
                </div>
                <div>
                    Hi Ali
                </div>
            </div>
        </div>
        <div>
            <div class="wrapper-container" id="wc"></div>
        </div>
    </div>
    <?php $this->load->view('template/mandatory_script', NULL, FALSE); ?>
    <script>
        var formatRupiah = (value) => {
            return value.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ".");
        }
        $(function() {
            $.ajax({
                url: `<?= $_ENV['API_URL']; ?>/products`,
                type: 'GET',
                beforeSend: function(xhr) {
                    xhr.setRequestHeader("X-API-Key", "wx9bHXTUDo");
                },
                success: function(response) {
                    console.log(response);
                    if (response.code == 200) {
                        for (let i = 0; i < response.data.length; i++) {
                            const element = response.data[i];
                            $('#wc').append(`<div class="card mb-3 card-product" onclick="document.location.href = 'App/detail/${element.id}';">
                                <div class="card-body p-2">
                                    <div class="row">
                                        <div class="col-5">
                                            <img class="object-fit-cover" src="${element.image}" style="width: 100%; height: 125px; border-radius: 5px; border: 1px solid #ccc;">
                                        </div>
                                        <div class="col-7 d-flex flex-column justify-content-between" style="margin-left: -12.5px;">
                                            <div>
                                                <h1 style="font-size: 18px;">${element.name}</h1>
                                                <p style="font-size: 12px; margin-top: -5px !important;" class="text-muted pb-0 mb-1">
                                                    ${element.description}
                                                </p>
                                            </div>
                                            <span>${element.price}</span>
                                        </div>
                                    </div>
                                </div>
                            </div>`);
                        }
                    }
                },
                error: function(xhr, error, code) {
                    if (xhr?.responseJSON?.error) {
                        errorAlert(xhr?.responseJSON?.error, ".wrapper");
                        return;
                    }
                    if (xhr?.responseJSON?.error?.user_detail_message) {
                        errorAlert(xhr?.responseJSON?.error?.user_detail_message, ".wrapper");
                        return;
                    }
                    if (xhr?.responseJSON?.message) {
                        errorAlert(xhr?.responseJSON?.message, ".wrapper");
                        return;
                    }
                    errorAlert(`${error}, ${(code == "" ? "failed to call API, please try again later." : code)}`, ".wrapper");
                    return;
                }
            });
        });
    </script>
</body>

</html>
