<!DOCTYPE html>
<html lang="en">

<head>
    <?php $this->load->view('template/head/meta', ['title' => 'Home', 'url' => site_url()], FALSE); ?>
    <?php $this->load->view('template/head/mandatory_style', NULL, FALSE); ?>
    <style>
        .nutrition-card {
            border: 1px solid #ddd;
            border-radius: 10px;
            box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
            padding: 15px;
        }

        .nutrition-table {
            width: 100%;
            font-size: 14px;
        }

        .nutrition-table td {
            vertical-align: top;
            padding: 2.5px 0;
        }

        .nutrition-table td:first-child {
            text-align: left;
        }

        .nutrition-table td:last-child {
            text-align: right;
        }

        .nutrition-header {
            background-color: #f9f9f9;
            border-top-left-radius: 10px;
            border-top-right-radius: 10px;
            padding: 10px 15px;
            color: #333;
            font-weight: bold;
        }

        .nutrition-table {
            font-family: monospace;
        }

        .nutri-score-grade-card {
            border: 1px solid #ddd;
            border-radius: 10px;
            box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
            position: relative;
        }

        .nutri-score-grade-header {
            background-color: #02813D;
            color: white;
            padding: 15px;
            border-top-left-radius: 10px;
            border-top-right-radius: 10px;
            display: flex;
            justify-content: space-between;
            align-items: center;
        }

        .nutri-score-grade-header h5 {
            margin: 0;
            font-weight: bold;
        }

        .nutri-score-grade-header span {
            font-size: 36px;
            /* font-weight: bold; */
            color: white;
        }

        .nutri-score-grade-body {
            padding: 15px;
            border-bottom-left-radius: 10px;
            border-bottom-right-radius: 10px;
            position: relative;
        }

        .nutri-score-grade-body p {
            margin: 0;
            font-size: 16px;
        }

        .info-mark {
            position: absolute;
            right: 15px;
            top: 50%;
            transform: translateY(-50%);
            font-size: 24px;
            color: #ccc;
            cursor: pointer;
        }

        .ingredient-row {
            display: flex;
            align-items: center;
            gap: 10px;
            margin-bottom: 10px;
        }

        .ingredient-name {
            flex: 2;
        }

        .ingredient-quantity {
            flex: 1;
        }

        .btn-remove {
            padding: 6px 10px;
            font-size: 14px;
        }
    </style>
</head>

<body>
    <div class="wrapper" style="padding-bottom: 20px;">
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
                <form action="">
                    <div class="card text-bg-dark text-center mb-3">
                        <img src="https://alibaba-2025.oss-ap-southeast-5.aliyuncs.com/tempe-kukus.jpeg" style="width: 100%; max-height: 250px; object-fit: cover;" class="shadow-sm" alt="">
                    </div>
                    <div class="card shadow-sm mb-3">
                        <div class="card-body">
                            <div class="mb-3">
                                <label for="" class="form-label">Name</label>
                                <input type="text" class="form-control" name="" id="" aria-describedby="" value="Tempe Kukus">
                            </div>
                            <div class="mb-2">
                                <label for="" class="form-label">Description</label>
                                <textarea name="" class="form-control" id="" rows="3">Tempe kukus tanpa minyak, tinggi protein dan serat.</textarea>
                            </div>
                        </div>
                    </div>
                    <!-- <div class="card text-bg-dark text-center mb-2">GRADE DAN TOTAL SCORE</div> -->
                    <div class="card nutri-score-grade-card">
                        <!-- Header -->
                        <div class="nutri-score-grade-header">
                            <h5>Nutri-Score Grade</h5>
                            <span>A</span>
                        </div>
                    </div>
                    <!-- <div class="card text-bg-dark text-center mb-2">NUTRISI, KALORI, PROTEIN, LEMAK DLL (COLLAPSE)</div> -->
                     <div class="card nutrition-card mb-3" style="border-radius: 0;">
                        <table class="nutrition-table">
                            <tbody>
                                <tr>
                                    <td>Energy</td>
                                    <td align="right">1</td>
                                    <td>KJ</td>
                                </tr>
                                <tr>
                                    <td>Saturated Fatty</td>
                                    <td align="right">0</td>
                                    <td>g</td>
                                </tr>
                                <tr>
                                    <td>Sugar</td>
                                    <td align="right">1.5</td>
                                    <td>g</td>
                                </tr>
                                <tr>
                                    <td>Salt</td>
                                    <td align="right">0.003</td>
                                    <td>g</td>
                                </tr>
                                <tr>
                                    <td>Protein</td>
                                    <td align="right">5</td>
                                    <td>g</td>
                                </tr>
                                <tr>
                                    <td>Fibres</td>
                                    <td align="right">5</td>
                                    <td>g</td>
                                </tr>
                                <tr>
                                    <td>Fruit, Vegetables & Legumes</td>
                                    <td align="right">5</td>
                                    <td>%</td>
                                </tr>
                            </tbody>
                        </table>
                    </div>
                    <!-- <div class="card p-4 shadow-sm mb-3">
                        <h5 class="mb-3">Ingredients</h5>
                        <div id="ingredients-list">
                            <div class="ingredient-row">
                                <input type="text" class="form-control ingredient-name" placeholder="">
                                <input type="text" class="form-control ingredient-quantity" placeholder="">
                                <button class="btn btn-danger btn-remove"><i class="bi bi-trash"></i></button>
                            </div>
                        </div>
                        <button id="add-ingredient" type="button" class="btn btn-success mt-2">+</button>
                    </div> -->
                    <div class="card nutrition-card mb-3">
                        <table class="nutrition-table">
                            <tbody>
                                <tr>
                                    <td style="vertical-align: middle !important;">Halal</td>
                                    <td><i class="bi bi-check-circle text-success fs-5"></i></td>
                                </tr>
                                <tr>
                                    <td style="vertical-align: middle !important;">Vegan</td>
                                    <td><i class="bi bi-check-circle text-success fs-5"></i></td>
                                </tr>
                                <tr>
                                    <td style="vertical-align: middle !important;">Friendly for Diabetes</td>
                                    <td><i class="bi bi-check-circle text-success fs-5"></i></td>
                                </tr>
                                <tr>
                                    <td style="vertical-align: middle !important;">Friendly for Hypertension</td>
                                    <td><i class="bi bi-check-circle text-success fs-5"></i></td>
                                </tr>
                                <tr>
                                    <td style="vertical-align: middle !important;">Friendly for Cholesterol</td>
                                    <td><i class="bi bi-check-circle text-success fs-5"></i></td>
                                </tr>
                                <tr>
                                    <td style="vertical-align: middle !important;">Friendly for Gout</td>
                                    <td><i class="bi bi-check-circle text-success fs-5"></i></td>
                                </tr>
                                <tr>
                                    <td style="vertical-align: middle !important;">Friendly for Kidney</td>
                                    <td><i class="bi bi-check-circle text-success fs-5"></i></td>
                                </tr>
                                <tr>
                                    <td style="vertical-align: middle !important;">Friendly for Celiac / Gluten-Free</td>
                                    <td><i class="bi bi-check-circle text-success fs-5"></i></td>
                                </tr>
                                <tr>
                                    <td style="vertical-align: middle !important;">Friendly for Thyroid</td>
                                    <td><i class="bi bi-check-circle text-success fs-5"></i></td>
                                </tr>
                                <tr>
                                    <td style="vertical-align: middle !important;">Friendly for Obesity</td>
                                    <td><i class="bi bi-check-circle text-success fs-5"></i></td>
                                </tr>
                                <tr>
                                    <td style="vertical-align: middle !important;">Friendly for Digestive Issues</td>
                                    <td><i class="bi bi-check-circle text-success fs-5"></i></td>
                                </tr>
                            </tbody>
                        </table>
                    </div>
                    <div class="card nutri-score-grade-card mb-3">
                        <div class="nutri-score-grade-header" style="background-color: #198754;">
                            <h5 style="font-weight: 0;">Sustainability Packaging</h5>
                        </div>
                        <div class="nutri-score-grade-body">
                            <p>-</p>
                        </div>
                    </div>
                    <!-- <div class="btn btn-success w-100">UPDATE</div> -->
                </form>
            </div>
        </div>
    </div>
    <?php $this->load->view('template/mandatory_script', NULL, FALSE); ?>
    <script>
        $(function() {
            $('#add-ingredient').on('click', function () {
                const newRow = `
                    <div class="ingredient-row">
                        <input type="text" class="form-control ingredient-name" placeholder="">
                        <input type="text" class="form-control ingredient-quantity" placeholder="">
                        <button class="btn btn-danger btn-remove"><i class="bi bi-trash"></i></button>
                    </div>
                `;
                $('#ingredients-list').append(newRow);
            });

            $(document).on('click', '.btn-remove', function () {
                $(this).closest('.ingredient-row').remove();
            });
        });
    </script>
</body>

</html>
