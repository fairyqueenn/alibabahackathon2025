<!DOCTYPE html>
<html lang="en">

<head>
    <?php $this->load->view('template/head/meta', ['title' => 'Home', 'url' => site_url()], FALSE); ?>
    <?php $this->load->view('template/head/mandatory_style', NULL, FALSE); ?>
</head>

<body>
    <div class="wrapper">
        <div class="wrapper-header">
            <div class="d-flex justify-content-between">
                <div>
                    <a href="javascript:history.back();" class="link-unstyled">
                        <i class="bi bi-chevron-left"></i>
                    </a>
                </div>
                <div>
                    Hi Ali
                </div>
            </div>
        </div>
        <div>
            <div class="wrapper-container">
                <form id="form-create">
                    <div class="mb-3">
                        <label class="form-label">Name</label>
                        <input type="text" class="form-control" name="name" required>
                    </div>
                    <div class="mb-3">
                        <label class="form-label">Description</label>
                        <textarea class="form-control" rows="3" name="description" required></textarea>
                    </div>
                    <div class="mb-3">
                        <label class="form-label">Price</label>
                        <input type="number" class="form-control" name="price" required>
                    </div>
                    <div class="mb-3">
                        <label class="form-label">Picture</label>
                        <input type="file" class="form-control" name="file" required>
                    </div>
                </form>
            </div>
        </div>
        <div style="position: absolute; bottom: 20px; right: 20px;">
            <button id="btn-create" type="button" class="btn btn-md btn-primary block" data-bs-toggle="tooltip" title="Create">
                <i class="bi bi-plus"></i> Create
            </button>
        </div>
    </div>
    <?php $this->load->view('template/mandatory_script', NULL, FALSE); ?>
    <script>
        function submitForm() {
            $.ajax({
                url: `<?= $_ENV['API_URL']; ?>/products`,
                type: 'POST',
                data: new FormData($('#form-create')[0]),
                contentType: false,
                processData: false,
                beforeSend: function(xhr) {
                    xhr.setRequestHeader("X-API-Key", "wx9bHXTUDo");
                    $('#btn-create').css('cursor', 'wait');
                    $('#btn-create').prop('disabled', true);

                },
                success: function(res) {
                    console.log(res);
                    window.location.href = "detail";
                },
                error: function(xhr, error, code) {
                    console.log(xhr?.responseJSON?.error)
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
                },
                complete: function() {
                    $('#btn-create').css('cursor', 'pointer');
                    $('#btn-create').prop('disabled', false);
                    $('#form-create')[0].reset();
                }
            });
        }
        $(function() {
            $('#btn-create').click(function() {
               $('#form-create').submit();
            });

            $('#form-create').submit(function(e) {
                e.preventDefault();
                submitForm();
            });
        });
    </script>
</body>

</html>
