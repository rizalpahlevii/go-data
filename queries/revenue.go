package queries

import (
	"fmt"
	"main/database"
)

//
//-- --------------------------------------------------------
//
//--
//-- Table structure for table `customer12692`
//--
//
//CREATE TABLE `customer12692` (
//`customer_id` varchar(255) NOT NULL,
//`customer_unique_id` varchar(255) NOT NULL,
//`customer_zip_code_prefix` varchar(255) NOT NULL,
//`customer_city` varchar(255) NOT NULL,
//`customer_state` varchar(255) NOT NULL
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
//
//-- --------------------------------------------------------
//
//--
//-- Table structure for table `geolocation12692`
//--
//
//CREATE TABLE `geolocation12692` (
//`geolocation_zip_code_prefix` varchar(50) DEFAULT NULL,
//`geolocation_lat` varchar(50) DEFAULT NULL,
//`geolocation_lng` varchar(50) DEFAULT NULL,
//`geolocation_city` varchar(50) DEFAULT NULL,
//`geolocation_state` varchar(50) DEFAULT NULL
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
//
//-- --------------------------------------------------------
//
//--
//-- Table structure for table `location12692`
//--
//
//CREATE TABLE `location12692` (
//`state` varchar(255) DEFAULT NULL,
//`state_code` varchar(255) DEFAULT NULL,
//`city` varchar(255) DEFAULT NULL,
//`zip_code_start` varchar(255) DEFAULT NULL,
//`zip_code_end` varchar(255) DEFAULT NULL,
//`state_lat` varchar(255) DEFAULT NULL,
//`state_lng` varchar(255) DEFAULT NULL
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
//
//-- --------------------------------------------------------
//
//--
//-- Table structure for table `order12692`
//--
//
//CREATE TABLE `order12692` (
//`order_id` varchar(50) NOT NULL,
//`customer_id` varchar(50) NOT NULL,
//`order_status` varchar(15) NOT NULL,
//`order_purchase_timestamp` timestamp NOT NULL,
//`order_approved_at` timestamp NULL DEFAULT NULL,
//`order_delivered_carrier_date` timestamp NULL DEFAULT NULL,
//`order_delivered_customer_date` timestamp NULL DEFAULT NULL,
//`order_estimated_delivery_date` timestamp NOT NULL
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
//
//-- --------------------------------------------------------
//
//--
//-- Table structure for table `order_item12692`
//--
//
//CREATE TABLE `order_item12692` (
//`order_id` varchar(50) NOT NULL,
//`order_item_id` varchar(50) NOT NULL,
//`product_id` varchar(50) NOT NULL,
//`seller_id` varchar(50) NOT NULL,
//`shipping_limit_date` timestamp NOT NULL,
//`price` float NOT NULL,
//`freight_value` float NOT NULL
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
//
//-- --------------------------------------------------------
//
//--
//-- Table structure for table `order_payment12692`
//--
//
//CREATE TABLE `order_payment12692` (
//`order_id` varchar(50) DEFAULT NULL,
//`payment_sequential` int DEFAULT NULL,
//`payment_type` varchar(50) DEFAULT NULL,
//`payment_installments` int DEFAULT NULL,
//`payment_value` float DEFAULT NULL
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
//
//-- --------------------------------------------------------
//
//--
//-- Table structure for table `order_review12692`
//--
//
//CREATE TABLE `order_review12692` (
//`review_id` varchar(50) DEFAULT NULL,
//`order_id` varchar(50) DEFAULT NULL,
//`review_score` int DEFAULT NULL,
//`review_comment_title` varchar(255) DEFAULT NULL,
//`review_comment_message` varchar(255) DEFAULT NULL,
//`review_creation_date` datetime DEFAULT NULL,
//`review_answer_timestamp` timestamp NULL DEFAULT NULL
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
//
//-- --------------------------------------------------------
//
//--
//-- Table structure for table `product12692`
//--
//
//CREATE TABLE `product12692` (
//`product_id` varchar(255) NOT NULL,
//`product_category_name` varchar(255) DEFAULT NULL,
//`product_name_length` int DEFAULT NULL,
//`product_description_length` int DEFAULT NULL,
//`product_photos_qty` int DEFAULT NULL,
//`product_weight_g` int DEFAULT NULL,
//`product_length_cm` int DEFAULT NULL,
//`product_height_cm` int DEFAULT NULL,
//`product_width_cm` int DEFAULT NULL
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
//
//-- --------------------------------------------------------
//
//--
//-- Table structure for table `seller12692`
//--
//
//CREATE TABLE `seller12692` (
//`seller_id` varchar(50) NOT NULL,
//`seller_zip_code_prefix` varchar(7) NOT NULL,
//`seller_city` varchar(255) DEFAULT NULL,
//`seller_state` varchar(255) DEFAULT NULL
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
//
//-- --------------------------------------------------------
//
//--
//-- Table structure for table `translation12692`
//--
//
//CREATE TABLE `translation12692` (
//`product_category_name` varchar(255) NOT NULL,
//`product_category_name_english` varchar(255) NOT NULL
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
//
//--
//-- Indexes for dumped tables
//--
//
//--
//-- Indexes for table `customer12692`
//--
//ALTER TABLE `customer12692`
//ADD PRIMARY KEY (`customer_id`);
//
//--
//-- Indexes for table `order12692`
//--
//ALTER TABLE `order12692`
//ADD PRIMARY KEY (`order_id`);
//
//--
//-- Indexes for table `product12692`
//--
//ALTER TABLE `product12692`
//ADD PRIMARY KEY (`product_id`);
//
//--
//-- Indexes for table `seller12692`
//--
//ALTER TABLE `seller12692`
//ADD PRIMARY KEY (`seller_id`);
//COMMIT;
//
///*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
///*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
///*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

func GetRevenues() {
	//Displays the revenue (profit) earned every day based on sales every day by looking at the price paid base on the table above
	db := database.GetConnection()
	defer db.Close()

	script := `SELECT DATE_FORMAT(order12692.order_purchase_timestamp, '%Y-%m-%d') AS date, SUM(order_item12692.price) AS revenue FROM order12692 INNER JOIN order_item12692 ON order12692.order_id = order_item12692.order_id GROUP BY date ORDER BY date ASC`

	rows, err := db.Query(script)
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var orderPurchaseTimestamp string
		var orderItemPrice float64
		err = rows.Scan(&orderPurchaseTimestamp, &orderItemPrice)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(orderPurchaseTimestamp, orderItemPrice)
	}
}

func Question2() {
	db := database.GetConnection()
	defer db.Close()

	//Displays the volume of products transported from seller to customer (in kilograms),
	//assuming the product is always physically stored at the seller's location

	script := `SELECT s.seller_id, s.seller_zip_code_prefix, s.seller_city, s.seller_state,
       p.product_id, p.product_category_name,
       SUM(p.product_weight_g/1000) AS total_weight
FROM product12692 p
INNER JOIN order_item12692 oi ON p.product_id = oi.product_id 
INNER JOIN seller12692 s ON oi.seller_id = s.seller_id
GROUP BY s.seller_id, p.product_id
ORDER BY total_weight ASC`

	rows, err := db.Query(script)
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var sellerId string
		var sellerZipCodePrefix string
		var sellerCity string
		var sellerState string
		var productId string
		var productCategoryName string
		var totalWeight float64
		err = rows.Scan(&sellerId, &sellerZipCodePrefix, &sellerCity, &sellerState, &productId, &productCategoryName, &totalWeight)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(sellerId, sellerZipCodePrefix, sellerCity, sellerState, productId, productCategoryName, totalWeight)
	}
}
