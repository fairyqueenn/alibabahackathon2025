<!DOCTYPE html>
<html lang="en">

<head>
    <?php $this->load->view('template/head/meta', ['title' => 'Home', 'url' => site_url()], FALSE); ?>
    <?php $this->load->view('template/head/mandatory_style', NULL, FALSE); ?>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/select2/4.0.13/css/select2.min.css" integrity="sha512-nMNlpuaDPrqlEls3IX/Q56H36qvBASwb3ipuo3MxeWbsQB1881ox0cRv7UPTgBlriqoynt35KjEwgGUeUXIPnw==" crossorigin="anonymous" referrerpolicy="no-referrer" />
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
            <div class="wrapper-container">
                <div class="mb-3">
                    <select id="multiSelect" multiple class="form-control">
                        <option value="">Diabetes</option>
                        <option value="">Hypertension</option>
                        <option value="">Cholesterol</option>
                        <option value="">Gout</option>
                        <option value="">Kidney</option>
                    </select>
                </div>
                <div class="card">
                    <div class="card-body p-2">
                        <div class="row">
                            <div class="col-5">
                                <img class="object-fit-cover" src="<?= base_url('assets/img/martabak.jpeg'); ?>" style="width: 100%; height: 125px; border-radius: 5px; border: 1px solid #ccc;">
                            </div>
                            <div class="col-7 d-flex flex-column justify-content-between" style="margin-left: -12.5px;">
                                <div>
                                    <h1 style="font-size: 18px;">Martabak Manis</h1>
                                    <p style="font-size: 12px; margin-top: -5px !important;" class="text-muted pb-0 mb-1">
                                        Lorem ipsum dolor sit amet consectetur adipisicing elit. Facilis minima quod provident?
                                    </p>
                                </div>
                                <span>Rp1.000.000</span>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="card mt-3">
                    <div class="card-body p-2">
                        <div class="row">
                            <div class="col-5">
                                <img class="object-fit-cover" src="<?= base_url('assets/img/martabak.jpeg'); ?>" style="width: 100%; height: 125px; border-radius: 5px; border: 1px solid #ccc;">
                            </div>
                            <div class="col-7 d-flex flex-column justify-content-between" style="margin-left: -12.5px;">
                                <div>
                                    <h1 style="font-size: 18px;">Martabak Manis</h1>
                                    <p style="font-size: 12px; margin-top: -5px !important;" class="text-muted pb-0 mb-1">
                                        Lorem ipsum dolor sit amet consectetur adipisicing elit. Facilis minima quod provident?
                                    </p>
                                </div>
                                <span>Rp1.000.000</span>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <?php $this->load->view('template/mandatory_script', NULL, FALSE); ?>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/select2/4.0.13/js/select2.min.js" integrity="sha512-2ImtlRlf2VVmiGZsjm9bEyhjGW4dU7B6TNwh/hx/iSByxNENtj3WVE6o/9Lj4TJeVXPi4bnOIMXFIJJAeufa0A==" crossorigin="anonymous" referrerpolicy="no-referrer"></script>
    <script>
        $(function() {
            $('#multiSelect').select2({
                placeholder: "Do you have any illness?",
                allowClear: true
            });
        });
    </script>
</body>

</html>
