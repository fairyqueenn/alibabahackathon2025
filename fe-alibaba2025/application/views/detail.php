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
                        <img id="product-image" src="" style="width: 100%; max-height: 250px; object-fit: cover;" class="shadow-sm" alt="">
                    </div>
                    <div class="card shadow-sm mb-3">
                        <div class="card-body">
                            <div class="mb-3">
                                <label for="" class="form-label">Name</label>
                                <input type="text" class="form-control" readonly aria-describedby="" value="" id="product-name">
                            </div>
                            <div class="mb-2">
                                <label for="" class="form-label">Description</label>
                                <textarea readonly class="form-control" rows="3" id="product-description"></textarea>
                            </div>
                        </div>
                    </div>
                    <div class="card nutri-score-grade-card">
                        <div class="nutri-score-grade-header" id="header-score">
                            <h5>Nutri-Score Grade</h5>
                            <span id="score"></span>
                        </div>
                    </div>
                     <div class="card nutrition-card mb-3" style="border-radius: 0;">
                        <table class="nutrition-table">
                            <tbody>
                                <tr>
                                    <td>Energy</td>
                                    <td align="right" id="energy"></td>
                                    <td>KJ</td>
                                </tr>
                                <tr>
                                    <td>Saturated Fatty</td>
                                    <td align="right" id="saturated_fatty"></td>
                                    <td>g</td>
                                </tr>
                                <tr>
                                    <td>Sugar</td>
                                    <td align="right" id="sugar"></td>
                                    <td>g</td>
                                </tr>
                                <tr>
                                    <td>Salt</td>
                                    <td align="right" id="salt"></td>
                                    <td>g</td>
                                </tr>
                                <tr>
                                    <td>Protein</td>
                                    <td align="right" id="protein"></td>
                                    <td>g</td>
                                </tr>
                                <tr>
                                    <td>Fibres</td>
                                    <td align="right" id="fibres"></td>
                                    <td>g</td>
                                </tr>
                                <tr>
                                    <td>Fruit, Vegetables & Legumes</td>
                                    <td align="right" id="fruit_vegetable_legumes"></td>
                                    <td>%</td>
                                </tr>
                            </tbody>
                        </table>
                    </div>
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
                            <h5 style="font-weight: 0;">Packaging Information</h5>
                        </div>
                        <div class="nutri-score-grade-body" id="packaging-info"></div>
                    </div>
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

            var scoreColorMap = {
                "A": "#1CA91E",
                "B": "#97C32B",
                "C": "#FEF21B",
                "D": "#FFAD19",
                "E": "#FF1612"
            }

            var id = "<?= $id; ?>";
            $.ajax({
                url: `<?= $_ENV['API_URL']; ?>/products/${id}`,
                type: 'GET',
                beforeSend: function(xhr) {
                    xhr.setRequestHeader("X-API-Key", "wx9bHXTUDo");
                },
                success: function(response) {
                    if (response.code == 200) {
                        const product = response.data;
                        $('#product-image').attr('src', product.image);
                        $('#product-name').val(product.name);
                        $('#product-description').val(product.description);
                        $('#salt').html(product.ingredients.nutrition.salt);
                        $('#sugar').html(product.ingredients.nutrition.sugar);
                        $('#energy').html(product.ingredients.nutrition.energy);
                        $('#fibres').html(product.ingredients.nutrition.fibres);
                        $('#protein').html(product.ingredients.nutrition.protein);
                        $('#saturated_fatty').html(product.ingredients.nutrition.saturated_fatty);
                        $('#fruit_vegetable_legumes').html(product.ingredients.nutrition.fruit_vegetable_legumes);
                        $('#score').html(product.ingredients.nutri_score);
                        $('#header-score').css("background-color", scoreColorMap[product.ingredients.nutri_score]);
                        if (product.ingredients.nutri_score == "C")  {
                            $('#header-score').css("color", "black");
                        }
                        const url = 'http://8.215.3.103:8000/packaging_summaries';
                        const data = {
                            query: `biasanya ${product.name} di indonesia dibungkus pake apa dan di pemasaran online? apakah bungkusnya bisa bermanfaat kembali tolong rating suistainablity pacakagingnya`
                        };

                        $.ajax({
                            url: url,
                            type: 'POST',
                            contentType: 'application/json',
                            data: JSON.stringify(data),
                            success: function (response) {
                                $('#packaging-info').text(response.response);
                            },
                            error: function (xhr, status, error) {
                                $('#packaging-info').text('Error: ' + error + '\nStatus: ' + status);
                            }
                        });
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
