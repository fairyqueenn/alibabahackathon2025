<script src="<?= base_url("node_modules/jquery/dist/jquery.min.js"); ?>"></script>
<script src="<?= base_url("node_modules/bootstrap/dist/js/bootstrap.bundle.min.js"); ?>"></script>
<script src="<?= base_url('assets/js-snackbar/dist/js-snackbar.min.js'); ?>"></script>
<script>
    function successAlert(message, el) {
        SnackBar({
            position: "bc",
            container: el,
            width: "500px",
            status: "success",
            icon: "✓",
            message: message
        });
    }

    function errorAlert(message, el) {
        SnackBar({
            position: "bc",
            container: el,
            width: "500px",
            status: "danger",
            icon: "danger",
            message: message
        });
    }

    function formatRupiah(value) {
        let result = new Intl.NumberFormat('id-ID', {
            minimumFractionDigits: 0
        }).format(value);

        return `Rp${result},-`;
    }
</script>
