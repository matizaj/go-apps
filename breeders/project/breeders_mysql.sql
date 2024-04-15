-- MariaDB dump 10.19  Distrib 10.11.2-MariaDB, for osx10.18 (x86_64)
--
-- Host: 127.0.0.1    Database: breeders
-- ------------------------------------------------------
-- Server version	10.6.7-MariaDB-1:10.6.7+maria~focal

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `breeders`
--

DROP TABLE IF EXISTS `breeders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `breeders` (
                            `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `breeders`
--

LOCK TABLES `breeders` WRITE;
/*!40000 ALTER TABLE `breeders` DISABLE KEYS */;
/*!40000 ALTER TABLE `breeders` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `cat_breeds`
--

DROP TABLE IF EXISTS `cat_breeds`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cat_breeds` (
                              `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
                              `breed` varchar(255) DEFAULT NULL,
                              `weight_low_lbs` int(11) DEFAULT NULL,
                              `weight_high_lbs` int(11) DEFAULT NULL,
                              `lifespan` int(11) DEFAULT NULL,
                              `details` text DEFAULT NULL,
                              `alternate_names` varchar(255) DEFAULT NULL,
                              `geographic_origin` varchar(255) DEFAULT NULL,
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=68 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cat_breeds`
--

LOCK TABLES `cat_breeds` WRITE;
/*!40000 ALTER TABLE `cat_breeds` DISABLE KEYS */;
INSERT INTO `cat_breeds` VALUES
                             (1,'Abyssinian',7,10,14,'The Abyssinian is easy to care for, and a joy to have in your home. They’re affectionate cats and love both people and other animals.','','Egypt'),
                             (2,'Aegean',7,10,10,'Native to the Greek islands known as the Cyclades in the Aegean Sea, these are natural cats, meaning they developed without humans getting involved in their breeding. As a breed, Aegean Cats are rare, although they are numerous on their home islands. They are generally friendly toward people and can be excellent cats for families with children.','','Greece'),
                             (3,'American Bobtail',7,16,12,'American Bobtails are loving and incredibly intelligent cats possessing a distinctive wild appearance. They are extremely interactive cats that bond with their human family with great devotion.','','United States'),
                             (4,'American Curl',5,10,12,'Distinguished by truly unique ears that curl back in a graceful arc, offering an alert, perky, happily surprised expression, they cause people to break out into a big smile when viewing their first Curl. Curls are very people-oriented, faithful, affectionate soulmates, adjusting remarkably fast to other pets, children, and new situations.','','United States'),
                             (5,'American Shorthair',8,15,16,'The American Shorthair is known for its longevity, robust health, good looks, sweet personality, and amiability with children, dogs, and other pets.','Domestic Shorthair','United States'),
                             (6,'American Wirehair',8,15,15,'The American Wirehair tends to be a calm and tolerant cat who takes life as it comes. His favorite hobby is bird-watching from a sunny windowsill, and his hunting ability will stand you in good stead if insects enter the house.','','United States'),
                             (7,'Arabian Mau',8,16,12,'Arabian Mau cats are social and energetic. Due to their energy levels, these cats do best in homes where their owners will be able to provide them with plenty of playtime, attention and interaction from their owners. These kitties are friendly, intelligent, and adaptable, and will even get along well with other pets and children.','Alley cat','United Arab Emirates'),
                             (8,'Australian Mist',7,15,12,'The Australian Mist thrives on human companionship. Tolerant of even the youngest of children, these friendly felines enjoy playing games and being part of the hustle and bustle of a busy household. They make entertaining companions for people of all ages, and are happy to remain indoors between dusk and dawn or to be wholly indoor pets.','Spotted Mist','Australia'),
                             (9,'Balinese',4,10,13,'Balinese are curious, outgoing, intelligent cats with excellent communication skills. They are known for their chatty personalities and are always eager to tell you their views on life, love, and what you’ve served them for dinner. ','Long-haired Siamese','United States'),
                             (10,'Bambino',4,9,12,'The Bambino is a breed of cat that was created as a cross between the Sphynx and the Munchkin breeds. The Bambino cat has short legs, large upright ears, and is usually hairless. They love to be handled and cuddled up on the laps of their family members.','','United States'),
                             (11,'Bengal',6,12,13,'Bengals are a lot of fun to live with, but they\'re definitely not the cat for everyone, or for first-time cat owners. Extremely intelligent, curious and active, they demand a lot of interaction and woe betide the owner who doesn\'t provide it.','','United States'),
                             (12,'Birman',6,15,15,'The Birman is a docile, quiet cat who loves people and will follow them from room to room. Expect the Birman to want to be involved in what you’re doing. He communicates in a soft voice, mainly to remind you that perhaps it’s time for dinner or maybe for a nice cuddle on the sofa. He enjoys being held and will relax in your arms like a furry baby.','Sacred Birman, Sacred Cat Of Burma','France'),
                             (13,'Bombay',6,11,12,'The the golden eyes and the shiny black coa of the Bopmbay is absolutely striking. Likely to bond most with one family member, the Bombay will follow you from room to room and will almost always have something to say about what you are doing, loving attention and to be carried around, often on his caregiver\'s shoulder.','Small black Panther','United States'),
(14,'British Longhair',8,18,13,'The British Longhair is a very laid-back relaxed cat, often perceived to be very independent although they will enjoy the company of an equally relaxed and likeminded cat. They are an affectionate breed, but very much on their own terms and tend to prefer to choose to come and sit with their owners rather than being picked up.','','United Kingdom'),
(15,'British Shorthair',12,20,12,'The British Shorthair is a very pleasant cat to have as a companion, ans is easy going and placid. The British is a fiercely loyal, loving cat and will attach herself to every one of her family members. While loving to play, she doesn\'t need hourly attention. If she is in the mood to play, she will find someone and bring a toy to that person. The British also plays well by herself, and thus is a good companion for single people.','Highlander, Highland Straight, Britannica','United Kingdom'),
                             (16,'Burmese',6,12,16,'Burmese love being with people, playing with them, and keeping them entertained. They crave close physical contact and abhor an empty lap. They will follow their humans from room to room, and sleep in bed with them, preferably under the covers, cuddled as close as possible. At play, they will turn around to see if their human is watching and being entertained by their crazy antics.','','Burma'),
                             (17,'Burmilla',6,13,12,'The Burmilla is a fairly placid cat. She tends to be an easy cat to get along with, requiring minimal care. The Burmilla is affectionate and sweet and makes a good companion, the Burmilla is an ideal companion to while away a lonely evening. Loyal, devoted, and affectionate, this cat will stay by its owner, always keeping them company.','','United Kingdom'),
                             (18,'California Spangled',10,15,15,'Perhaps the only thing about the California spangled cat that isn’t wild-like is its personality. Known to be affectionate, gentle and sociable, this breed enjoys spending a great deal of time with its owners. They are very playful, often choosing to perch in high locations and show off their acrobatic skills.','Spangle','United States'),
                             (19,'Chantilly-Tiffany',7,12,15,'The Chantilly is a devoted companion and prefers company to being left alone. While the Chantilly is not demanding, she will \"chirp\" and \"talk\" as if having a conversation. This breed is affectionate, with a sweet temperament. It can stay still for extended periods, happily lounging in the lap of its loved one. This quality makes the Tiffany an ideal traveling companion, and an ideal house companion for senior citizens and the physically handicapped.','Chantilly, Foreign Longhair','United States'),
                             (20,'Chartreux',6,15,15,'The Chartreux is generally silent but communicative. Short play sessions, mixed with naps and meals are their perfect day. Whilst appreciating any attention you give them, they are not demanding, content instead to follow you around devotedly, sleep on your bed and snuggle with you if you’re not feeling well.','','France'),
                             (21,'Chausie',7,15,12,'For those owners who desire a feline capable of evoking the great outdoors, the strikingly beautiful Chausie retains a bit of the wild in its appearance but has the house manners of our friendly, familiar moggies. Very playful, this cat needs a large amount of space to be able to fully embrace its hunting instincts.','Nile Cat','Egypt'),
                             (22,'Cheetoh',8,15,12,'The Cheetoh has a super affectionate nature and real love for their human companions; they are intelligent with the ability to learn quickly. You can expect that a Cheetoh will be a fun-loving kitty who enjoys playing, running, and jumping through every room in your house.',' ','United States'),
                             (23,'Colorpoint Shorthair',4,10,13,'Colorpoint Shorthairs are an affectionate breed, devoted and loyal to their people. Sensitive to their owner’s moods, Colorpoints are more than happy to sit at your side or on your lap and purr words of encouragement on a bad day. They will constantly seek out your lap whenever it is open and in the moments when your lap is preoccupied they will stretch out in sunny spots on the ground.','','United States'),
                             (24,'Cornish Rex',5,9,12,'This is a confident cat who loves people and will follow them around, waiting for any opportunity to sit in a lap or give a kiss. He enjoys being handled, making it easy to take him to the veterinarian or train him for therapy work. The Cornish Rex stay in kitten mode most of their lives and well into their senior years. ','','United Kingdom'),
                             (25,'Cymric',8,13,13,'The Cymric is a placid, sweet cat. They do not get too upset about anything that happens in their world. They are loving companions and adore people. They are smart and dexterous, capable of using his paws to get into cabinets or to open doors.','Spangle','Canada'),
                             (26,'Cyprus',8,16,12,'Loving, loyal, social and inquisitive, the Cyprus cat forms strong ties with their families and love nothing more than to be involved in everything that goes on in their surroundings. They are not overly active by nature which makes them the perfect companion for people who would like to share their homes with a laid-back relaxed feline companion. ','Cypriot cat','Cyprus'),
                             (27,'Devon Rex',5,10,12,'The favourite perch of the Devon Rex is right at head level, on the shoulder of her favorite person. She takes a lively interest in everything that is going on and refuses to be left out of any activity. Count on her to stay as close to you as possible, occasionally communicating his opinions in a quiet voice. She loves people and welcomes the attentions of friends and family alike.','Pixie cat, Alien cat, Poodle cat','United Kingdom'),
                             (28,'Donskoy',10,12,13,'Donskoy are affectionate, intelligent, and easy-going. They demand lots of attention and interaction. The Donskoy also gets along well with other pets. It is now thought the same gene that causes degrees of hairlessness in the Donskoy also causes alterations in cat personality, making them calmer the less hair they have.','','Russia'),
                             (29,'Dragon Li',9,12,22,'The Dragon Li is loyal, but not particularly affectionate. They are known to be very intelligent, and their natural breed status means that they\'re very active. She is is gentle with people, and has a reputation as a talented hunter of rats and other vermin.','Chinese Lia Hua, Li Hua','China'),
(30,'Egyptian Mau',6,14,19,'The Egyptian Mau is gentle and reserved. She loves her people and desires attention and affection from them but is wary of others. Early, continuing socialization is essential with this sensitive and sometimes shy cat, especially if you plan to show or travel with her. Otherwise, she can be easily startled by unexpected noises or events.','Pharaoh Cat','Egypt'),
(31,'European Burmese',7,14,12,'The European Burmese is a very affectionate, intelligent, and loyal cat. They thrive on companionship and will want to be with you, participating in everything you do. While they might pick a favorite family member, chances are that they will interact with everyone in the home, as well as any visitors that come to call. They are inquisitive and playful, even as adults. ','','Burma'),
(32,'Exotic Shorthair',7,14,12,'The Exotic Shorthair is a gentle friendly cat that has the same personality as the Persian. They love having fun, don’t mind the company of other cats and dogs, also love to curl up for a sleep in a safe place. Exotics love their own people, but around strangers they are cautious at first. Given time, they usually warm up to visitors.','Exotic','United States'),
(33,'Havana Brown',6,10,12,'The Havana Brown is human oriented, playful, and curious. She has a strong desire to spend time with her people and involve herself in everything they do. Being naturally inquisitive, the Havana Brown reaches out with a paw to touch and feel when investigating curiosities in its environment. They are truly sensitive by nature and frequently gently touch their human companions as if they are extending a paw of friendship.','Havana, HB','United Kingdom'),
(34,'Himalayan',7,14,13,'Calm and devoted, Himalayans make excellent companions, though they prefer a quieter home. They are playful in a sedate kind of way and enjoy having an assortment of toys. The Himalayan will stretch out next to you, sleep in your bed and even sit on your lap when she is in the mood.','Himalayan Persian, Colourpoint Persian, Longhaired Colourpoint, Himmy','United States'),
(35,'Japanese Bobtail',5,10,15,'The Japanese Bobtail is an active, sweet, loving and highly intelligent breed. They love to be with people and play seemingly endlessly. They learn their name and respond to it. They bring toys to people and play fetch with a favorite toy for hours. Bobtails are social and are at their best when in the company of people. They take over the house and are not intimidated. If a dog is in the house, Bobtails assume Bobtails are in charge.','Japanese Truncated Cat','Japan'),
(36,'Javanese',5,10,14,'Javanese are endlessly interested, intelligent and active. They tend to enjoy jumping to great heights, playing with fishing pole-type or other interactive toys and just generally investigating their surroundings. He will attempt to copy things you do, such as opening doors or drawers.',' ','United States'),
(37,'Khao Manee',8,12,12,'The Khao Manee is highly intelligent, with an extrovert and inquisitive nature, however they are also very calm and relaxed, making them an idea lap cat.','Diamond Eye cat','Thailand'),
(38,'Korat',7,11,12,'The Korat is a natural breed, and one of the oldest stable cat breeds. They are highly intelligent and confident cats that can be fearless, although they are startled by loud sounds and sudden movements. Korats form strong bonds with their people and like to cuddle and stay nearby.','','Thailand'),
(39,'Kurilian',8,15,17,'The character of the Kurilian Bobtail is independent, highly intelligent, clever, inquisitive, sociable, playful, trainable, absent of aggression and very gentle. They are devoted to their humans and when allowed are either on the lap of or sleeping in bed with their owners.','','Russia'),
(40,'LaPerm',6,10,14,'LaPerms are gentle and affectionate but also very active. Unlike many active breeds, the LaPerm is also quite content to be a lap cat. The LaPerm will often follow your lead; that is, if they are busy playing and you decide to sit and relax, simply pick up your LaPerm and sit down with it, and it will stay in your lap, devouring the attention you give it.','Si-Sawat','Thailand'),
(41,'Maine Coon',12,18,13,'They are known for their size and luxurious long coat Maine Coons are considered a gentle giant. The good-natured and affable Maine Coon adapts well to many lifestyles and personalities. She likes being with people and has the habit of following them around, but isn’t needy. Most Maine Coons love water and they can be quite good swimmers.','Coon Cat, Maine Cat, Maine Shag, Snowshoe Cat, American Longhair, The Gentle Giants','United States'),
(42,'Malayan',6,3,14,'Malayans love to explore and even enjoy traveling by way of a cat carrier. They are quite a talkative and rather loud cat with an apparent strong will. These cats will make sure that you give it the attention it seeks and always seem to want to be held and hugged. They will constantly interact with people, even strangers. They love to play and cuddle.','Asian','United Kingdom'),
(43,'Manx',7,13,12,'The Manx is a placid, sweet cat that is gentle and playful. She never seems to get too upset about anything. She is a loving companion and adores being with people.','Manks, Stubbin, Rumpy','Isle of Man'),
(44,'Munchkin',5,9,14,'The Munchkin is an outgoing cat who enjoys being handled. She has lots of energy and is faster and more agile than she looks. The shortness of their legs does not seem to interfere with their running and leaping abilities.','','United States'),
(45,'Nebelung',7,11,12,'The Nebelung may have a reserved nature, but she loves to play (being especially fond of retrieving) and enjoys jumping or climbing to high places where she can study people and situations at her leisure before making up her mind about whether she wants to get involved.','Longhaired Russian Blue','United States'),
(46,'Norwegian Forest Cat',8,16,15,'The Norwegian Forest Cat is a sweet, loving cat. She appreciates praise and loves to interact with her parent. She makes a loving companion and bonds with her parents once she accepts them for her own. She is still a hunter at heart. She loves to chase toys as if they are real. She is territorial and patrols several times each day to make certain that all is fine.','Skogkatt / Skaukatt, Norsk Skogkatt / Norsk Skaukatt, Weegie','Norway'),
(47,'Ocicat',7,15,15,'Loyal and devoted to their owners, the Ocicat is intelligent, confident, outgoing, and seems to have many dog traits. They can be trained to fetch toys, walk on a lead, taught to \'speak\', come when called, and follow other commands. ','','United States'),
(48,'Oriental',5,10,14,'Orientals are passionate about the people in their lives. They become extremely attached to their humans, so be prepared for a lifetime commitment. When you are not available to entertain her, an Oriental will divert herself by jumping on top of the refrigerator, opening drawers, seeking out new hideaways.','Foreign Type','United States'),
(49,'Persian',9,14,14,'Persians are sweet, gentle cats that can be playful or quiet and laid-back. Great with families and children, they absolutely love to lounge around the house. While they don’t mind a full house or active kids, they’ll usually hide when they need some alone time.','Longhair, Persian Longhair, Shiraz, Shirazi','Iran (Persia)'),
(50,'Pixie-bob',8,17,15,'Companionable and affectionate, the Pixie-bob wants to be an integral part of the family. The Pixie-Bob’s ability to bond with their humans along with their patient personas make them excellent companions for children.','','United States'),
(51,'Ragamuffin',8,20,14,'The Ragamuffin is calm, even tempered and gets along well with all family members. Changes in routine generally do not upset her. She is an ideal companion for those in apartments, and with children due to her patient nature.','','United States'),
(52,'Ragdoll',12,20,15,'Ragdolls love their people, greeting them at the door, following them around the house, and leaping into a lap or snuggling in bed whenever given the chance. They are the epitome of a lap cat, enjoy being carried and collapsing into the arms of anyone who holds them.','Rag doll','United States'),
(53,'Russian Blue',5,11,12,'Russian Blues are very loving and reserved. They do not like noisy households but they do like to play and can be quite active when outdoors. They bond very closely with their owner and are known to be compatible with other pets.','Archangel Blue, Archangel Cat','Russia'),
(54,'Savannah',8,25,18,'Savannah is the feline version of a dog. Actively seeking social interaction, they are given to pouting if left out. Remaining kitten-like through life. Profoundly loyal to immediate family members whilst questioning the presence of strangers. Making excellent companions that are loyal, intelligent and eager to be involved.','','United States'),
(55,'Scottish Fold',5,11,18,'The Scottish Fold is a sweet, charming breed. She is an easy cat to live with and to care for. She is affectionate and is comfortable with all members of her family. Her tail should be handled gently. Folds are known for sleeping on their backs, and for sitting with their legs stretched out and their paws on their belly. This is called the \"Buddha Position\".','Scot Fold','United Kingdom'),
(56,'Selkirk Rex',6,16,15,'The Selkirk Rex is an incredibly patient, loving, and tolerant breed. The Selkirk also has a silly side and is sometimes described as clownish. She loves being a lap cat and will be happy to chat with you in a quiet voice if you talk to her. ','Shepherd Cat','United States'),
(57,'Siamese',8,15,12,'While Siamese cats are extremely fond of their people, they will follow you around and supervise your every move, being talkative and opinionated. They are a demanding and social cat, that do not like being left alone for long periods.','Siam, Thai Cat','Thailand'),
(58,'Siberian',8,16,15,'The Siberians dog like temperament and affection makes the ideal lap cat and will live quite happily indoors. Very agile and powerful, the Siberian cat can easily leap and reach high places, including the tops of refrigerators and even doors. ','Moscow Semi-longhair, HairSiberian Forest Cat','Russia'),
(59,'Singapura',5,8,12,'The Singapura is usually cautious when it comes to meeting new people, but loves attention from his family so much that she sometimes has the reputation of being a pest. This is a highly active, curious and affectionate cat. She may be small, but she knows she’s in charge','Drain Cat, Kucinta, Pura','Singapore'),
(60,'Snowshoe',7,12,17,'The Snowshoe is a vibrant, energetic, affectionate and intelligent cat. They love being around people which makes them ideal for families, and becomes unhappy when left alone for long periods of time. Usually attaching themselves to one person, they do whatever they can to get your attention.','','United States'),
(61,'Somali',6,12,12,'The Somali lives life to the fullest. He climbs higher, jumps farther, plays harder. Nothing escapes the notice of this highly intelligent and inquisitive cat. Somalis love the company of humans and other animals.','Fox Cat, Long-Haired Abyssinian','Somalia'),
(62,'Sphynx',6,12,12,'The Sphynx is an intelligent, inquisitive, extremely friendly people-oriented breed. Sphynx commonly greet their owners  at the front door, with obvious excitement and happiness. She has an unexpected sense of humor that is often at odds with her dour expression.','Canadian Hairless, Canadian Sphynx','Canada'),
(63,'Tonkinese',6,12,15,'Intelligent and generous with their affection, a Tonkinese will supervise all activities with curiosity. Loving, social, active, playful, yet content to be a lap cat','Tonk','Canada'),
(64,'Toyger',7,15,14,'The Toyger has a sweet, calm personality and is generally friendly. He\'s outgoing enough to walk on a leash, energetic enough to play fetch and other interactive games, and confident enough to get along with other cats and friendly dogs.','','United States'),
(65,'Turkish Angora',5,10,17,'This is a smart and intelligent cat which bonds well with humans. With its affectionate and playful personality the Angora is a top choice for families. The Angora gets along great with other pets in the home, but it will make clear who is in charge, and who the house belongs to','Ankara','Turkey'),
(66,'Turkish Van',7,20,12,'While the Turkish Van loves to jump and climb, play with toys, retrieve and play chase, she is is big and ungainly; this is one cat who doesn’t always land on his feet. While not much of a lap cat, the Van will be happy to cuddle next to you and sleep in your bed. ','Turkish Cat, Swimming cat','Turkey'),
(67,'York Chocolate',12,18,14,'York Chocolate cats are known to be true lap cats with a sweet temperament. They love to be cuddled and petted. Their curious nature makes them follow you all the time and participate in almost everything you do, even if it\'s related to water: unlike many other cats, York Chocolates love it.','York','United States');
/*!40000 ALTER TABLE `cat_breeds` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `cats`
--

DROP TABLE IF EXISTS `cats`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cats` (
                        `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
                        `cat_name` varchar(255) NOT NULL,
                        `breed_id` int(11) unsigned DEFAULT NULL,
                        `breeder_id` int(11) unsigned DEFAULT NULL,
                        `color` varchar(255) NOT NULL,
                        `date_of_birth` date NOT NULL,
                        `spayed_neutered` int(11) NOT NULL DEFAULT 0,
                        `description` text NOT NULL,
                        `weight` int(11) NOT NULL,
                        PRIMARY KEY (`id`),
                        KEY `breed_id` (`breed_id`),
                        KEY `breeder_id` (`breeder_id`),
                        CONSTRAINT `cats_ibfk_1` FOREIGN KEY (`breed_id`) REFERENCES `cat_breeds` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
                        CONSTRAINT `cats_ibfk_2` FOREIGN KEY (`breeder_id`) REFERENCES `breeders` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cats`
--

LOCK TABLES `cats` WRITE;
/*!40000 ALTER TABLE `cats` DISABLE KEYS */;
/*!40000 ALTER TABLE `cats` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `dog_breeds`
--

DROP TABLE IF EXISTS `dog_breeds`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `dog_breeds` (
                              `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                              `breed` varchar(255) DEFAULT NULL,
                              `weight_low_lbs` int(11) DEFAULT NULL,
                              `weight_high_lbs` int(11) DEFAULT NULL,
                              `lifespan` int(11) DEFAULT NULL,
                              `details` text DEFAULT NULL,
                              `alternate_names` varchar(255) DEFAULT NULL,
                              `geographic_origin` varchar(255) DEFAULT NULL,
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=579 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `dog_breeds`
--

LOCK TABLES `dog_breeds` WRITE;
/*!40000 ALTER TABLE `dog_breeds` DISABLE KEYS */;
INSERT INTO `dog_breeds` VALUES
                             (390,'Affenpinscher',8,12,14,'','',''),
                             (391,'Afghan Hound',50,60,14,'','',''),
                             (392,'Aidi',50,55,13,'','Atlas Mountain Dog',''),
                             (393,'Airedale Terrier',45,45,13,'','',''),
                             (394,'Akbash Dog',75,140,10,'Possesses a unique combination of Mastiff and gazehound characteristics (UKC).','',''),
                             (395,'Akita',80,120,12,'','',''),
                             (396,'Alaskan Husky',35,60,13,'The Alaskan Husky is not considered a \"purebred\" dog like the Siberian Husky. Alaskan huskies are bred for speed primarily for dog sledding. Their genetic makeup generally includes a mix of Northern breeds including the Siberian Husky','',''),
                             (397,'Alaskan Klee Kai',16,23,15,'','',''),
                             (398,'Alaskan Malamute',75,124,12,'','',''),
                             (399,'American Bulldog',60,125,11,'','',''),
                             (400,'American Bully',35,120,11,'The American Bully breed developed as a natural extension of the American Pit Bull Terrier. (UKC)','','United States'),
                             (401,'American Cocker Spaniel',15,30,12,'','',''),
                             (402,'American Eskimo Dog',25,30,14,'','',''),
                             (403,'American Foxhound',65,70,12,'','',''),
                             (404,'American Hairless Terrier',12,16,15,'','','United States'),
                             (405,'American Leopard Hound',45,70,14,'','Leopard Hound','United States'),
                             (406,'American Pit Bull Terrier',30,75,11,'','','United States'),
                             (407,'American Staffordshire Terrier',40,70,14,'','','United States'),
                             (408,'American Water Spaniel',25,45,12,'','',''),
                             (409,'Anatolian Shepherd Dog',100,150,12,'','',''),
                             (410,'Australian Cattle Dog',33,50,14,'','Queensland Heeler',''),
                             (411,'Australian Kelpie',33,50,12,'','',''),
                             (412,'Australian Shepherd',40,65,14,'','',''),
                             (413,'Australian Terrier',15,20,13,'','',''),
                             (414,'Barbet',35,65,13,'','','France'),
                             (415,'Basset Hound',33,70,13,'','',''),
                             (416,'Beagle',18,35,13,'','',''),
                             (417,'Bearded Collie',40,60,13,'','',''),
                             (418,'Beauceron',65,110,11,'','\"Berger de Beauce',''),
                             (419,'Bedlington Terrier',17,23,14,'','',''),
                             (420,'Belgian Sheepdog-Malinois',40,80,15,'','Belgian Malinois',''),
                             (421,'Belgian Sheepdog-Tervuren',45,75,13,'','Belgian Tervuren',''),
                             (422,'Bernese Mountain Dog',70,115,9,'','',''),
                             (423,'Bichon Frise',7,18,15,'','',''),
                             (424,'Black & Tan Coonhound',40,110,11,'','American Black and Tan Coonhound',''),
                             (425,'Black Mouth Cur',35,100,14,'','',''),
                             (426,'Black Russian Terrier',80,143,11,'','',''),
                             (427,'Bloodhound',79,119,11,'','',''),
                             (428,'Bluetick Coonhound',45,80,12,'','',''),
                             (429,'Border Collie',27,55,14,'','',''),
                             (430,'Border Terrier',11,16,14,'','',''),
                             (431,'Borzoi',55,105,12,'','',''),
                             (432,'Boston Terrier',12,25,12,'','Boston Bulldog',''),
                             (433,'Bouvier Des Flandres',59,110,11,'','',''),
                             (434,'Boykin Spaniel',25,40,13,'','','United States'),
                             (435,'Boxer',55,80,11,'','',''),
                             (436,'Briard',55,100,12,'','\"Berger de Brie',''),
                             (437,'Brittany',30,45,13,'','Brittany Spaniel',''),
                             (438,'Brussels Griffon',7,13,14,'','',''),
                             (439,'Bull Terrier',25,80,13,'','English Bull Terrier',''),
                             (440,'Bulldog',40,58,9,'','\"British Bulldog',''),
                             (441,'Bullmastiff',90,130,8,'','',''),
                             (442,'Cairn Terrier',13,17,14,'','',''),
                             (443,'Canaan Dog',35,55,14,'','Bedouin Sheepdog',''),
                             (444,'Cane Corso',80,110,11,'','',''),
                             (445,'Cardigan Welsh Corgi',25,38,14,'','\"Cardi','Welsh Cardigan Corgi \"'),
                             (446,'Catahoula Leopard Dog',50,95,12,'','Louisiana Catahoula Leopared Dog',''),
                             (447,'Cavalier King Charles Spaniel',11,20,14,'','',''),
                             (448,'Chesapeake Bay Retriever',55,80,12,'','\"Chesapeake Bay Ducking Dog',''),
                             (449,'Dalmatian',45,71,12,'','',''),
                             (450,'Chihuahua',1,7,15,'','',''),
                             (451,'Chinese Crested',5,13,16,'','',''),
                             (452,'Chinese Shar-Pei',40,60,10,'','',''),
                             (453,'Chow Chow',45,70,10,'','',''),
                             (454,'Clumber Spaniel',55,85,11,'','',''),
                             (455,'Collie',45,75,13,'','',''),
                             (456,'Coton de Tulear',8,15,17,'','',''),
                             (457,'Curly Coated Retriever',60,95,11,'','',''),
                             (458,'Dachshund',1,11,14,'','',''),
                             (459,'Dandie Dinmont Terrier',18,24,14,'','',''),
                             (460,'Danish-Swedish Farmdog',15,20,12,'','',''),
                             (461,'Doberman Pinscher',60,100,11,'','',''),
                             (462,'Dogo Argentino',80,100,12,'','\"Argentino Dogo',''),
                             (463,'Dutch Shepherd',42,75,13,'','Holland Shepherd',''),
                             (464,'English Bulldog',40,58,9,'','British Bulldog',''),
                             (465,'English Cocker Spaniel',24,35,13,'','',''),
                             (466,'English Coonhound',45,70,12,'','\"American Coonhound',''),
                             (467,'English Foxhound',45,65,12,'','',''),
                             (468,'English Setter',45,80,12,'','',''),
                             (469,'English Springer Spaniel',35,50,13,'','',''),
                             (470,'English Toy Spaniel',8,14,11,'','',''),
                             (471,'Field Spaniel',35,50,13,'','',''),
                             (472,'Finnish Spitz',15,35,14,'','',''),
                             (473,'Flat Coated Retriever',60,80,9,'','',''),
                             (474,'French Bulldog',16,28,11,'','',''),
                             (475,'German Longhaired Pointer',55,80,13,'','',''),
                             (476,'German Pinscher',25,45,13,'','',''),
                             (477,'German Shepherd Dog',48,90,13,'','',''),
                             (478,'German Shorthaired Pointer',45,70,11,'','',''),
                             (479,'German Wirehaired Pointer',50,70,15,'','',''),
                             (480,'Giant Schnauzer',55,104,14,'','',''),
                             (481,'Glen of Imaal Terrier',32,40,13,'','',''),
                             (482,'Golden Retriever',55,75,11,'','',''),
                             (483,'Gordon Setter',45,80,13,'','',''),
                             (484,'Great Dane',99,200,9,'','',''),
                             (485,'Great Pyrenees',85,115,13,'','',''),
                             (486,'Greater Swiss Mountain Dog',85,140,10,'','',''),
                             (487,'Greyhound',60,70,12,'','English Greyhound',''),
                             (488,'Harrier',40,65,14,'','',''),
                             (489,'Havanese',7,14,15,'','Havana Silk Dog',''),
                             (490,'Ibizan Hound',40,55,13,'','',''),
                             (491,'Irish Setter',50,75,14,'','Irish Red Setter ',''),
                             (492,'Irish Terrier',24,28,14,'','',''),
                             (493,'Irish Water Spaniel',45,68,13,'','',''),
                             (494,'Irish Wolfhound',89,125,7,'','',''),
                             (495,'Italian Greyhound',7,14,15,'','',''),
                             (496,'Jack Russell Terrier',9,18,13,'','Russell Terrier',''),
                             (497,'Japanese Chin',4,15,11,'','',''),
                             (498,'Kangal',90,150,14,'','\"Kangal Shepherd Dog',''),
                             (499,'Keeshond',35,66,14,'','',''),
                             (500,'Kerry Blue Terrier',33,45,14,'','Irish Blue Terrier',''),
                             (501,'Komondor',80,135,11,'','',''),
                             (502,'Kuvasz',70,115,11,'','',''),
                             (503,'Labrador Retriever',55,80,11,'','',''),
                             (504,'Lakeland Terrier',15,18,14,'','',''),
                             (505,'Lhasa Apso',12,18,14,'','Tibetan Apso ',''),
                             (506,'Maltese',4,9,14,'','',''),
                             (507,'Manchester Terrier (Standard)',12,22,16,'','',''),
                             (508,'Manchester Terrier (Toy)',6,12,16,'','',''),
                             (509,'Mastiff',120,230,8,'','\"English Mastiff',''),
                             (510,'Miniature Bull Terrier',18,28,12,'','',''),
                             (511,'Miniature Pinscher',8,13,14,'','',''),
                             (512,'Neapolitan Mastiff',110,154,8,'','Italian Mastiff',''),
                             (513,'Nederlandse Kooikerhondje',20,30,14,'FCI Classification: Group 8 Retrivers - Flushing Dogs - Water Dogs','',''),
                             (514,'Newfoundland',100,152,10,'','',''),
                             (515,'Norfolk Terrier',10,12,14,'','',''),
                             (516,'Norrbottenspets',20,30,16,'FCI Classification: Group 5 Spitz and primitive types Section 2 Nordic hunting dogs','',''),
                             (517,'Norwegian Elkhound',48,55,14,'','',''),
                             (518,'Nova Scotia Duck Tolling Retriever',35,51,13,'','',''),
                             (519,'Old English Sheepdog',60,100,11,'','\"Bob','72'),
                             (520,'Olde English Bulldogge',45,85,13,'','',''),
                             (521,'Otterhound',80,125,12,'','',''),
                             (522,'Papillon',5,11,15,'','',''),
                             (523,'Parson Russell Terrier',13,17,14,'','',''),
                             (524,'Patterdale Terrier',10,17,12,'','Black Fell Terrier',''),
                             (525,'Pekingese',6,14,13,'','','China'),
                             (526,'Pembroke Welsh Corgi',20,30,13,'','','Wales'),
                             (527,'Perro de Presa Canario',84,143,10,'','',''),
                             (528,'Petit Basset Griffon Vendeen',30,45,13,'','',''),
                             (529,'Pharaoh Hound',45,55,13,'','Kelb Tal-Fenek ','Malta'),
                             (530,'Plott Hound',40,75,13,'','','United States '),
                             (531,'Pointer',44,75,15,'','English Pointer','United Kingdom'),
                             (532,'Polish Lowland Sheepdog',30,50,13,'FCI Classification: Group 1 Sheepdogs and Cattle Dogs Section 1 Sheepdogs','Polish Owczarek Nizinny','Poland'),
                             (533,'Pomeranian',3,7,14,'','Toy German Spitz','Germany'),
                             (534,'Poodle Toy',4,8,14,'','','Germany/France'),
                             (535,'Poodle',10,39,14,'','',''),
                             (536,'Portuguese Water Dog',35,60,12,'','','Portugal'),
                             (537,'Pug',14,18,14,'','\"Carlin',''),
                             (538,'Puli',22,35,13,'','Hungarian Puli','Hungary'),
                             (539,'Pumi',17,33,13,'The Pumi came into being during the 17th to the 18th century in Hungary by cross breeding the primitive Puli with imported German and French Terrier type breeds with prick ears. It has been recognized as an independent breed since the beginning of the 20th century.','','Hungary'),
                             (540,'Rat Terrier',10,35,15,'','','United States '),
                             (541,'Redbone Coonhound',45,70,14,'','','United States '),
                             (542,'Rhodesian Ridgeback',65,85,12,'','African Lion Dog','41'),
                             (543,'Rottweiler',90,135,10,'','','Germany'),
                             (544,'Saint Bernard',120,200,9,'','Alpine Mastiff','Switzerland '),
                             (545,'Saluki',40,65,14,'Among the world\'s oldest breeds the slim but rugged Saluki was the hunting hound of kings for thousands of years. Salukis are swift and agile sprinters who love a good chase. ','',''),
(546,'Samoyed',35,65,13,'These dogs were kept by a native tribe to ancient Siberia known as the Samoyeds. The Samoyeds were used to pull sleds and herd reindeer.','',''),
(547,'Schapendoes',26,55,14,'','Dutch Sheepdog','Netherlands'),
(548,'Schipperke',10,19,13,'Well-suited for closed environments, the Schipperke is energetic and playful.','','Moorke'),
(549,'Schnauzer',28,50,15,'','',''),
(550,'Scottish Deerhound',75,110,10,'','','Scotland'),
(551,'Scottish Terrier',18,23,13,'','Aberdeen Terrier','Great Britain'),
(552,'Sealyham Terrier',18,25,13,'The Sealyham Terrier was born of a desire to produce the perfect terrier. ','Welsh Corgi',''),
(553,'Shetland Sheepdog',14,20,13,'','Sheltie','Japan'),
(554,'Shiba Inu',16,24,15,'','','Tibet/China'),
(555,'Shih Tzu',9,18,14,'','Chinese Lion Dog','20'),
(556,'Siberian Husky',35,60,13,'','Arctic Husky','Australia'),
(557,'Silky Terrier',8,10,14,'','Australian Silky Terrier',''),
(558,'Skye Terrier',18,40,13,'','','England'),
(559,'Smooth Fox Terrier',15,18,14,'','','Ireland'),
(560,'Spinone Italiano',62,86,11,'','Italian Coarse-Haired Pointer',''),
(561,'Staffordshire Bull Terrier',24,38,13,'','','England'),
(562,'Sussex Spaniel',35,51,14,'','','Germany'),
(563,'Tibetan Mastiff',70,160,11,'','Do-Khyi ','England'),
(564,'Tibetan Spaniel',9,15,14,'','','Tibet  '),
(565,'Tibetan Terrier',18,30,16,'','Dhoki Apso ','Tibet'),
(566,'Toy Fox Terrier',4,8,14,'','American Toy Terrier','Tibet'),
(567,'Treeing Tennessee Brindle',30,50,11,'','Treeing Walker',''),
(568,'Treeing Walker Coonhound',50,70,13,'','','United States'),
(569,'Vizsla',44,66,13,'','','United States'),
(570,'Weimaraner',55,90,12,'','','Hungary'),
(571,'Welsh Springer Spaniel',35,55,14,'','','Germany'),
(572,'Welsh Terrier',19,22,14,'','','Wales'),
(573,'West Highland White Terrier',15,22,14,'','Westie','Wales'),
(574,'Wheaten Terrier',30,40,13,'','','53'),
(575,'Whippet',25,40,14,'','','Great Britain'),
(576,'Wire Fox Terrier',15,18,14,'','','England'),
(577,'Wirehaired Pointing Griffon',35,70,14,'A relatively young breed.','',''),
(578,'Yorkshire Terrier',4,7,13,'','Broken-Haired Scotch Terrier','10');
/*!40000 ALTER TABLE `dog_breeds` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `dogs`
--

DROP TABLE IF EXISTS `dogs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `dogs` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `dog_name` varchar(255) NOT NULL,
  `breed_id` int(11) unsigned DEFAULT NULL,
  `breeder_id` int(11) unsigned DEFAULT NULL,
  `color` varchar(255) NOT NULL,
  `date_of_birth` date NOT NULL,
  `spayed_neutered` int(11) NOT NULL DEFAULT 0,
  `description` text NOT NULL,
  `weight` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `breed_id` (`breed_id`),
  KEY `breeder_id` (`breeder_id`),
  CONSTRAINT `dogs_ibfk_1` FOREIGN KEY (`breed_id`) REFERENCES `dog_breeds` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
  CONSTRAINT `dogs_ibfk_2` FOREIGN KEY (`breeder_id`) REFERENCES `breeders` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `dogs`
--

LOCK TABLES `dogs` WRITE;
/*!40000 ALTER TABLE `dogs` DISABLE KEYS */;
/*!40000 ALTER TABLE `dogs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `first_name` varchar(255) NOT NULL,
  `last_name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(60) NOT NULL,
  `user_active` int(11) NOT NULL DEFAULT 0,
  `access_level` int(11) NOT NULL DEFAULT 10,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES
(1,'Admin','User','admin@example.com','$2a$14$lfQ071jRtaUB6lNXorl7mOjxIlNbla9MWnQJwnZz2n2PM8ML2Velu',1,30);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-03-20 14:41:49
