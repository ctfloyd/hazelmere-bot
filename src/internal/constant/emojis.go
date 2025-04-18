package constant

import "github.com/ctfloyd/hazelmere-api/src/pkg/api"

var Emojis = map[api.ActivityType]string{
	api.ActivityTypeUnknown:                      "<:unknown:1358096257624834411>",
	api.ActivityTypeOverall:                      "<:overall:1358095954242306058>",
	api.ActivityTypeAttack:                       "<:attack:1358095033915408634>",
	api.ActivityTypeDefence:                      "<:defence:1358095384861479032>",
	api.ActivityTypeStrength:                     "<:strength:1358096108894814431>",
	api.ActivityTypeHitpoints:                    "<:hitpoints:1358095640088936674>",
	api.ActivityTypeRanged:                       "<:ranged:1358096013478596690>",
	api.ActivityTypePrayer:                       "<:prayer:1358095999549308958>",
	api.ActivityTypeMagic:                        "<:magic:1358095881639038976>",
	api.ActivityTypeCooking:                      "<:cooking:1358095267601322116>",
	api.ActivityTypeWoodcutting:                  "<:woodcutting:1358096340491567361>",
	api.ActivityTypeFletching:                    "<:fletching:1358095505812361293>",
	api.ActivityTypeFishing:                      "<:fishing:1358095495666598200>",
	api.ActivityTypeFiremaking:                   "<:firemaking:1358095488666173734>",
	api.ActivityTypeCrafting:                     "<:crafting:1358095292456505484>",
	api.ActivityTypeSmithing:                     "<:smithing:1358096081736700034>",
	api.ActivityTypeMining:                       "<:mining:1358103199017468116>",
	api.ActivityTypeHerblore:                     "<:herblore:1358095569935138897>",
	api.ActivityTypeAgility:                      "<:agility:1358095004979040526>",
	api.ActivityTypeThieving:                     "<:thieving:1358096209436217654>",
	api.ActivityTypeSlayer:                       "<:slayer:1358096075495309564>",
	api.ActivityTypeFarming:                      "<:farming:1358095475508773068>",
	api.ActivityTypeRunecraft:                    "<:runecrafting:1358096028783611962>",
	api.ActivityTypeHunter:                       "<:hunter:1358095657109553213>",
	api.ActivityTypeConstruction:                 "<:construction:1358095255202697236>",
	api.ActivityTypeLeaguePoints:                 "<:league_points:1358095868590428335>",
	api.ActivityTypeBountyHunterHunter:           "<:bounty_hunter_hunter:1358095052995301621>",
	api.ActivityTypeBountyHunterRogue:            "<:bounty_hunter_rogue:1358095060226412687>",
	api.ActivityTypeClueScrollsall:               "<:clue_scrolls_all:1358095129776230571>",
	api.ActivityTypeClueScrollsbeginner:          "<:clue_scrolls_beginner:1358095136705478696>",
	api.ActivityTypeClueScrollseasy:              "<:clue_scrolls_easy:1358095143571292380>",
	api.ActivityTypeClueScrollsmedium:            "<:clue_scrolls_medium:1358095201758871582>",
	api.ActivityTypeClueScrollshard:              "<:clue_scrolls_hard:1358095164391948428>",
	api.ActivityTypeClueScrollselite:             "<:clue_scrolls_elite:1358095155407622161>",
	api.ActivityTypeClueScrollsmaster:            "<:clue_scrolls_master:1358095187867598848>",
	api.ActivityTypeLMSRank:                      "<:last_man_standing:1358095862625992825>",
	api.ActivityTypePvPArenaRank:                 "<:pvp_arena:1358096006721573048>",
	api.ActivityTypeSoulWarsZeal:                 "<:soul_wars_zeal:1358096095628234792>",
	api.ActivityTypeRiftsclosed:                  "<:guardians_of_the_rift:1358095532878205164>",
	api.ActivityTypeColosseumGlory:               "<:colosseum_glory:1358095215507804431>",
	api.ActivityTypeAbyssalSire:                  "<:abyssal_sire:1358094997349601551>",
	api.ActivityTypeAlchemicalHydra:              "<:alchemical_hydra:1358095011824271580>",
	api.ActivityTypeArtio:                        "<:artio:1358095026428842124>",
	api.ActivityTypeBarrowsChests:                "<:barrows_chests:1358095045940609264>",
	api.ActivityTypeBryophyta:                    "<:bryophyta:1358095067239284766>",
	api.ActivityTypeCallisto:                     "<:callisto:1358095074344570900>",
	api.ActivityTypeCalvarion:                    "<:calvarion:1358095083768910116>",
	api.ActivityTypeCerberus:                     "<:cerberus:1358095089787994262>",
	api.ActivityTypeChambersofXeric:              "<:chambers_of_xeric:1358095100022099968>",
	api.ActivityTypeChambersofXericChallengeMode: "<:chambers_of_xeric_challenge_mode:1358095109081792833>",
	api.ActivityTypeChaosElemental:               "<:chaos_elemental:1358095116253921320>",
	api.ActivityTypeChaosFanatic:                 "<:chaos_fanatic:1358095122755227794>",
	api.ActivityTypeCommanderZilyana:             "<:commander_zilyana:1358095243177623725>",
	api.ActivityTypeCorporealBeast:               "<:corporeal_beast:1358095279693500596>",
	api.ActivityTypeCrazyArchaeologist:           "<:crazy_archaeologist:1358095324102529306>",
	api.ActivityTypeDagannothPrime:               "<:dagannoth_prime:1358095334840205470>",
	api.ActivityTypeDagannothRex:                 "<:dagannoth_rex:1358095351365632153>",
	api.ActivityTypeDagannothSupreme:             "<:dagannoth_supreme:1358095368956543198>",
	api.ActivityTypeDerangedArchaeologist:        "<:deranged_archaeologist:1358095395217211582>",
	api.ActivityTypeDukeSucellus:                 "<:duke_sucellus:1358095459981201519>",
	api.ActivityTypeGeneralGraardor:              "<:general_graardor:1358095512917508228>",
	api.ActivityTypeGiantMole:                    "<:giant_mole:1358095520354009088>",
	api.ActivityTypeGrotesqueGuardians:           "<:grotesque_guardians:1358095526142410912>",
	api.ActivityTypeHespori:                      "<:hespori:1358095599773421688>",
	api.ActivityTypeKalphiteQueen:                "<:kalphite_queen:1358095725602275450>",
	api.ActivityTypeKingBlackDragon:              "<:king_black_dragon:1358095835275198604>",
	api.ActivityTypeKraken:                       "<:kraken:1358095842895990916>",
	api.ActivityTypeKreeArra:                     "<:kreearra:1358095850303258854>",
	api.ActivityTypeKrilTsutsaroth:               "<:kril_tsutsaroth:1358095856619880468>",
	api.ActivityTypeLunarChests:                  "<:lunar_chests:1358095875372617978>",
	api.ActivityTypeMimic:                        "<:mimic:1358095888655843510>",
	api.ActivityTypeNex:                          "<:nex:1358095902006444213>",
	api.ActivityTypeNightmare:                    "<:nightmare:1358095933350613063>",
	api.ActivityTypePhosanisNightmare:            "<:phosanis_nightmare:1358095968960254135>",
	api.ActivityTypeObor:                         "<:obor:1358095944436154608>",
	api.ActivityTypePhantomMuspah:                "<:phantom_muspah:1358095961724948720>",
	api.ActivityTypeSarachnis:                    "<:sarachnis:1358096036492480643>",
	api.ActivityTypeScorpia:                      "<:scorpia:1358096053043204437>",
	api.ActivityTypeScurrius:                     "<:scurrius:1358096061205319701>",
	api.ActivityTypeSkotizo:                      "<:skotizo:1358096068277178410>",
	api.ActivityTypeSolHeredit:                   "<:sol_heredit:1358096089072533626>",
	api.ActivityTypeSpindel:                      "<:spindel:1358096102469009580>",
	api.ActivityTypeTempoross:                    "<:tempoross:1358096115819610252>",
	api.ActivityTypeTheGauntlet:                  "<:the_gauntlet:1358096147012391116>",
	api.ActivityTypeTheCorruptedGauntlet:         "<:the_corrupted_gauntlet:1358096139576016976>",
	api.ActivityTypeTheLeviathan:                 "<:the_leviathan:1358096155317113034>",
	api.ActivityTypeTheWhisperer:                 "<:the_whisperer:1358096169691250749>",
	api.ActivityTypeTheatreOfBlood:               "<:theatre_of_blood:1358096181296758914>",
	api.ActivityTypeTheatreOfBloodHardMode:       "<:theatre_of_blood_hard_mode:1358096193607045331>",
	api.ActivityTypeThermonuclearSmokeDevil:      "<:thermonuclear_smoke_devil:1358096201370570772>",
	api.ActivityTypeTombsOfAmascut:               "<:tombs_of_amascut:1358096217065783476>",
	api.ActivityTypeTombsOfAmascutExpertMode:     "<:tombs_of_amascut_expert:1358096223466291300>",
	api.ActivityTypeTzKalZuk:                     "<:tzkal_zuk:1358096230944735512>",
	api.ActivityTypeTzTokJad:                     "<:tztok_jad:1358096237194252339>",
	api.ActivityTypeVardorvis:                    "<:vardorvis:1358096267338579978>",
	api.ActivityTypeVenenatis:                    "<:venenatis:1358096277488795892>",
	api.ActivityTypeVetion:                       "<:vetion:1358096289178587336>",
	api.ActivityTypeVorkath:                      "<:vorkath:1358096301073633331>",
	api.ActivityTypeWintertodt:                   "<:wintertodt:1358096331582996683>",
	api.ActivityTypeZalcano:                      "<:zalcano:1358096348930379776>",
	api.ActivityTypeZulrah:                       "<:zulrah:1358094750359748670>",
	api.ActivityTypeAmoxliatl:                    "<:amoxliatl:1362831287462920514>",
	api.ActivityTypeAraxxor:                      "<:araxxor:1362831310871200179>",
	api.ActivityTypeTheHueycoatl:                 "<:the_hueycoatl:1362831427464335673>",
	api.ActivityTypeTheRoyalTitans:               "<:the_royal_titans:1362831445609025787>",
	api.ActivityTypeCollectionsLogged:            "<:collections_logged:1362831484901134336>",
}
