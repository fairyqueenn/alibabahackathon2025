<?php $us1 = $this->uri->segment(1); ?>
<nav class="navbar fixed-bottom bg-white menu-bottom">
    <div class="w-full">
        <center>
            <div class="row">
                <div class="d-grid col-6 text-center">
                    <a class="btn btn-left<?= ($us1 == "home" || $us1 == "") ? ' fw-bolder disabled' : ''; ?>" href="<?= site_url('home'); ?>"<?= ($us1 == "home" || $us1 == "") ? ' style="color: var(--vr-color-dark);"' : ''; ?>>
                        <div class="icon-section">
                            <i class="bi bi-house-door"></i>
                        </div>
                        <div class="text-section">
                            Home
                        </div>
                    </a>
                </div>
                <div class="d-grid col-6 text-center">
                    <a class="btn btn-right<?= $us1 == "activity" ? ' fw-bolder disabled' : ''; ?>" href="<?= site_url('activity'); ?>"<?= $us1 == "activity" ? ' style="color: var(--vr-color-dark);"' : ''; ?>>
                        <div class="icon-section">
                            <i class="bi bi-activity"></i>
                        </div>
                        <div class="text-section">
                            Activity
                        </div>
                    </a>
                </div>
            </div>
        </center>
    </div>
</nav>
