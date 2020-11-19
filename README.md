# megafile
MegaFile (.MEG) reader.

File format information from official Electronic Arts Command & Conquer Remastered Collection source code at https://github.com/electronicarts/CnC_Remastered_Collection


1. Buy Command & Conquer™ Remastered Collection
1. Extract files
1. 🎶🤘😎🤘🎶

# Extract files:

```shell
% ./megafile -x ./steamapps/common/CnCRemastered/Data/MUSIC.MEG 
file #001(idx:193)/203: data/audio/music/tdr_mus_reaching_out.wav 5760586 bytes at offset 13016
file #002(idx:13)/203: data/audio/music/rab_mus_the_plan_am.wav 9472126 bytes at offset 5773604
file #003(idx:32)/203: data/audio/music/rac_mus_hell_march_retaliation_remix.wav 2555916 bytes at offset 15245732
file #004(idx:149)/203: data/audio/music/tdc_mus_rain_in_the_night.wav 1859206 bytes at offset 17801648
file #005(idx:30)/203: data/audio/music/rac_mus_groundwire.wav 2771726 bytes at offset 19660856
file #006(idx:96)/203: data/audio/music/rar_mus_traction.wav 11647726 bytes at offset 22432584
file #007(idx:34)/203: data/audio/music/rac_mus_map_theme.wav 785966 bytes at offset 34080312
file #008(idx:14)/203: data/audio/music/rab_mus_the_truth_is_a_lie.wav 10244506 bytes at offset 34866280
file #009(idx:48)/203: data/audio/music/rac_mus_terminate.wav 3831106 bytes at offset 45110788
file #010(idx:171)/203: data/audio/music/tdr_mus_drone.wav 13126266 bytes at offset 48941896
file #011(idx:154)/203: data/audio/music/tdc_mus_target_mechanical_man.wav 2057726 bytes at offset 62068164
file #012(idx:158)/203: data/audio/music/tdc_mus_warfare_full_stop.wav 2164406 bytes at offset 64125892
file #013(idx:99)/203: data/audio/music/rar_mus_twin_cannon_retaliation_remix.wav 10973766 bytes at offset 66290300
file #014(idx:155)/203: data/audio/music/tdc_mus_to_be_feared.wav 1950836 bytes at offset 77264068
file #015(idx:139)/203: data/audio/music/tdc_mus_in_trouble.wav 2266396 bytes at offset 79214904
file #016(idx:173)/203: data/audio/music/tdr_mus_gdi_map_theme.wav 2965626 bytes at offset 81481300
file #017(idx:97)/203: data/audio/music/rar_mus_trenches.wav 15458246 bytes at offset 84446928
file #018(idx:82)/203: data/audio/music/rar_mus_mud_retaliation_remix.wav 13854406 bytes at offset 99905176
file #019(idx:175)/203: data/audio/music/tdr_mus_great_shot_extended.wav 2741206 bytes at offset 113759584
file #020(idx:56)/203: data/audio/music/rac_mus_vector.wav 3109546 bytes at offset 116500792
file #021(idx:87)/203: data/audio/music/rar_mus_roll_out.wav 11333986 bytes at offset 119610340
file #022(idx:153)/203: data/audio/music/tdc_mus_take_em_out.wav 2169516 bytes at offset 130944328
file #023(idx:27)/203: data/audio/music/rac_mus_floating.wav 3327946 bytes at offset 133113844
file #024(idx:68)/203: data/audio/music/rar_mus_crush_retaliation_remix.wav 10737726 bytes at offset 136441792
file #025(idx:151)/203: data/audio/music/tdc_mus_recon.wav 3113886 bytes at offset 147179520
file #026(idx:156)/203: data/audio/music/tdc_mus_untamed_land.wav 2217186 bytes at offset 150293408
file #027(idx:128)/203: data/audio/music/tdc_mus_depth_charge.wav 3052636 bytes at offset 152510596
file #028(idx:170)/203: data/audio/music/tdr_mus_drill.wav 12888406 bytes at offset 155563232
file #029(idx:66)/203: data/audio/music/rar_mus_chaos.wav 12409886 bytes at offset 168451640
file #030(idx:88)/203: data/audio/music/rar_mus_run.wav 15097466 bytes at offset 180861528
file #031(idx:84)/203: data/audio/music/rar_mus_radio_2.wav 11828466 bytes at offset 195958996
file #032(idx:121)/203: data/audio/music/tdc_mus_act_on_instinct.wav 1999276 bytes at offset 207787464
file #033(idx:44)/203: data/audio/music/rac_mus_running_through_pipes.wav 3330256 bytes at offset 209786740
file #034(idx:86)/203: data/audio/music/rar_mus_reload_fire.wav 14089046 bytes at offset 213116996
file #035(idx:118)/203: data/audio/music/tdb_mus_tank_battle_co.wav 9253446 bytes at offset 227206044
file #036(idx:162)/203: data/audio/music/tdr_mus_airstrike.wav 9545346 bytes at offset 236459492
file #037(idx:50)/203: data/audio/music/rac_mus_the_second_hand.wav 3306386 bytes at offset 246004840
file #038(idx:192)/203: data/audio/music/tdr_mus_rain_in_the_night.wav 7448846 bytes at offset 249311228
file #039(idx:92)/203: data/audio/music/rar_mus_snake.wav 13674226 bytes at offset 256760076
file #040(idx:186)/203: data/audio/music/tdr_mus_nod_score.wav 1637586 bytes at offset 270434304
file #041(idx:94)/203: data/audio/music/rar_mus_the_search.wav 13595126 bytes at offset 272071892
file #042(idx:76)/203: data/audio/music/rar_mus_hell_march.wav 18531806 bytes at offset 285667020
file #043(idx:141)/203: data/audio/music/tdc_mus_just_do_it_up.wav 1689596 bytes at offset 304198828
file #044(idx:67)/203: data/audio/music/rar_mus_crush.wav 11068266 bytes at offset 305888424
file #045(idx:77)/203: data/audio/music/rar_mus_hell_march_retaliation_remix.wav 10223226 bytes at offset 316956692
file #046(idx:201)/203: data/audio/music/tdr_mus_warfare_full_stop.wav 8747906 bytes at offset 327179920
file #047(idx:102)/203: data/audio/music/rar_mus_voice_rhythm_2.wav 12536586 bytes at offset 335927828
file #048(idx:108)/203: data/audio/music/tdb_mus_command_and_conquer_fkts.wav 8870126 bytes at offset 348464416
file #049(idx:105)/203: data/audio/music/rar_mus_workmen_retaliation_remix.wav 9912566 bytes at offset 357334544
file #050(idx:43)/203: data/audio/music/rac_mus_run.wav 3845596 bytes at offset 367247112
file #051(idx:203)/203: data/audio/music/tdr_mus_we_will_stop_them_deception_ost_version.wav 9153766 bytes at offset 371092708
file #052(idx:143)/203: data/audio/music/tdc_mus_nod_map_theme.wav 444086 bytes at offset 380246476
file #053(idx:24)/203: data/audio/music/rac_mus_dense.wav 3630976 bytes at offset 380690564
file #054(idx:59)/203: data/audio/music/rac_mus_workmen.wav 3414116 bytes at offset 384321540
file #055(idx:180)/203: data/audio/music/tdr_mus_in_trouble_ost_version.wav 10240586 bytes at offset 387735656
file #056(idx:9)/203: data/audio/music/rab_mus_hm_2_3_medley_fkts.wav 12157746 bytes at offset 397976244
file #057(idx:5)/203: data/audio/music/rab_mus_crush_fkts.wav 8613506 bytes at offset 410133992
file #058(idx:125)/203: data/audio/music/tdc_mus_cc_thang.wav 2298876 bytes at offset 418747500
file #059(idx:190)/203: data/audio/music/tdr_mus_prepare_for_battle.wav 10087566 bytes at offset 421046376
file #060(idx:33)/203: data/audio/music/rac_mus_intro_menu.wav 2521546 bytes at offset 431133944
file #061(idx:4)/203: data/audio/music/rab_mus_concrete_am.wav 11480566 bytes at offset 433655492
file #062(idx:1)/203: data/audio/music/rab_mus_big_foot_fkts.wav 16253866 bytes at offset 445136060
file #063(idx:46)/203: data/audio/music/rac_mus_smash.wav 3351046 bytes at offset 461389928
file #064(idx:18)/203: data/audio/music/rac_mus_backstab.wav 3422866 bytes at offset 464740976
file #065(idx:119)/203: data/audio/music/tdb_mus_target_mechanical_man_fkts.wav 8783886 bytes at offset 468163844
file #066(idx:147)/203: data/audio/music/tdc_mus_prepare_for_battle.wav 2502716 bytes at offset 476947732
file #067(idx:103)/203: data/audio/music/rar_mus_wasteland.wav 11867666 bytes at offset 479450448
file #068(idx:136)/203: data/audio/music/tdc_mus_heartbreak.wav 2444546 bytes at offset 491318116
file #069(idx:195)/203: data/audio/music/tdr_mus_ride_of_the_valkyries.wav 14546286 bytes at offset 493762664
file #070(idx:101)/203: data/audio/music/rar_mus_vector.wav 12459446 bytes at offset 508308952
file #071(idx:174)/203: data/audio/music/tdr_mus_great_shot.wav 1975826 bytes at offset 520768400
file #072(idx:177)/203: data/audio/music/tdr_mus_industrial.wav 8368506 bytes at offset 522744228
file #073(idx:64)/203: data/audio/music/rar_mus_big_foot.wav 15203586 bytes at offset 531112736
file #074(idx:6)/203: data/audio/music/rab_mus_grinder_1_2_medley_fkts.wav 8380686 bytes at offset 546316324
file #075(idx:120)/203: data/audio/music/tdb_mus_warfare_full_stop_fkts.wav 9298386 bytes at offset 554697012
file #076(idx:16)/203: data/audio/music/rac_mus_arazoid.wav 3167506 bytes at offset 563995400
file #077(idx:127)/203: data/audio/music/tdc_mus_demolition.wav 2142356 bytes at offset 567162908
file #078(idx:80)/203: data/audio/music/rar_mus_militant_force.wav 5311046 bytes at offset 569305264
file #079(idx:138)/203: data/audio/music/tdc_mus_in_the_line_of_fire.wav 1479106 bytes at offset 574616312
file #080(idx:146)/203: data/audio/music/tdc_mus_on_the_prowl.wav 2165526 bytes at offset 576095420
file #081(idx:112)/203: data/audio/music/tdb_mus_just_do_it_up_fkts.wav 7103466 bytes at offset 578260948
file #082(idx:100)/203: data/audio/music/rar_mus_underlying_thoughts.wav 12225366 bytes at offset 585364416
file #083(idx:85)/203: data/audio/music/rar_mus_radio_2_retaliation_remix.wav 10411666 bytes at offset 597589784
file #084(idx:181)/203: data/audio/music/tdr_mus_iron_fist.wav 10151826 bytes at offset 608001452
file #085(idx:54)/203: data/audio/music/rac_mus_twin_cannon_retaliation_remix.wav 2732246 bytes at offset 618153280
file #086(idx:164)/203: data/audio/music/tdr_mus_cc_80s_mix.wav 11728226 bytes at offset 620885528
file #087(idx:145)/203: data/audio/music/tdc_mus_no_mercy.wav 2423406 bytes at offset 632613756
file #088(idx:191)/203: data/audio/music/tdr_mus_radio.wav 8746646 bytes at offset 635037164
file #089(idx:202)/203: data/audio/music/tdr_mus_we_will_stop_them_deception.wav 9138926 bytes at offset 643783812
file #090(idx:157)/203: data/audio/music/tdc_mus_voice_rhythm.wav 3651206 bytes at offset 652922740
file #091(idx:117)/203: data/audio/music/tdb_mus_slave_to_the_system_fkts.wav 7930726 bytes at offset 656573948
file #092(idx:166)/203: data/audio/music/tdr_mus_creeping_upon.wav 10508826 bytes at offset 664504676
file #093(idx:45)/203: data/audio/music/rac_mus_shut_it.wav 3222246 bytes at offset 675013504
file #094(idx:188)/203: data/audio/music/tdr_mus_no_mercy_ost_version.wav 9714606 bytes at offset 678235752
file #095(idx:74)/203: data/audio/music/rar_mus_gloom.wav 11655566 bytes at offset 687950360
file #096(idx:161)/203: data/audio/music/tdr_mus_act_on_instinct_ost_version.wav 8341486 bytes at offset 699605928
file #097(idx:53)/203: data/audio/music/rac_mus_twin_cannon.wav 2828846 bytes at offset 707947416
file #098(idx:113)/203: data/audio/music/tdb_mus_mad_rap_fkts.wav 10528706 bytes at offset 710776264
file #099(idx:142)/203: data/audio/music/tdc_mus_march_to_doom.wav 1859276 bytes at offset 721304972
file #100(idx:163)/203: data/audio/music/tdr_mus_canyon_chase.wav 7826286 bytes at offset 723164248
file #101(idx:17)/203: data/audio/music/rac_mus_awaiting.wav 3201596 bytes at offset 730990536
file #102(idx:75)/203: data/audio/music/rar_mus_groundwire.wav 11309486 bytes at offset 734192132
file #103(idx:178)/203: data/audio/music/tdr_mus_in_the_line_of_fire.wav 6026166 bytes at offset 745501620
file #104(idx:29)/203: data/audio/music/rac_mus_gloom.wav 2867836 bytes at offset 751527788
file #105(idx:133)/203: data/audio/music/tdc_mus_gdi_map_theme.wav 725906 bytes at offset 754395624
file #106(idx:111)/203: data/audio/music/tdb_mus_industrial_fkts.wav 8416666 bytes at offset 755121532
file #107(idx:20)/203: data/audio/music/rac_mus_bog.wav 2570686 bytes at offset 763538200
file #108(idx:169)/203: data/audio/music/tdr_mus_die.wav 7810886 bytes at offset 766108888
file #109(idx:165)/203: data/audio/music/tdr_mus_cc_thang.wav 9304266 bytes at offset 773919776
file #110(idx:55)/203: data/audio/music/rac_mus_underlying_thoughts.wav 3029256 bytes at offset 783224044
file #111(idx:148)/203: data/audio/music/tdc_mus_radio.wav 2167206 bytes at offset 786253300
file #112(idx:159)/203: data/audio/music/tdc_mus_we_will_stop_them_deception.wav 2261076 bytes at offset 788420508
file #113(idx:63)/203: data/audio/music/rar_mus_backstab.wav 13889126 bytes at offset 790681584
file #114(idx:179)/203: data/audio/music/tdr_mus_in_trouble.wav 9212846 bytes at offset 804570712
file #115(idx:73)/203: data/audio/music/rar_mus_fogger.wav 14873046 bytes at offset 813783560
file #116(idx:58)/203: data/audio/music/rac_mus_wasteland.wav 2931046 bytes at offset 828656608
file #117(idx:19)/203: data/audio/music/rac_mus_big_foot.wav 3784766 bytes at offset 831587656
file #118(idx:168)/203: data/audio/music/tdr_mus_depth_charge.wav 12299846 bytes at offset 835372424
file #119(idx:199)/203: data/audio/music/tdr_mus_untamed_land.wav 8966586 bytes at offset 847672272
file #120(idx:36)/203: data/audio/music/rac_mus_mud.wav 3451216 bytes at offset 856638860
file #121(idx:35)/203: data/audio/music/rac_mus_militant_force.wav 1311106 bytes at offset 860090076
file #122(idx:2)/203: data/audio/music/rab_mus_blow_it_up_fkts.wav 9225726 bytes at offset 861401184
file #123(idx:38)/203: data/audio/music/rac_mus_no_mercy_retaliation_remix.wav 2563336 bytes at offset 870626912
file #124(idx:110)/203: data/audio/music/tdb_mus_got_a_present_for_ya_fkts.wav 6877786 bytes at offset 873190248
file #125(idx:98)/203: data/audio/music/rar_mus_twin_cannon.wav 11364786 bytes at offset 880068036
file #126(idx:116)/203: data/audio/music/tdb_mus_rain_in_the_night_fkts.wav 10833626 bytes at offset 891432824
file #127(idx:182)/203: data/audio/music/tdr_mus_just_do_it_up.wav 6832006 bytes at offset 902266452
file #128(idx:7)/203: data/audio/music/rab_mus_hell_march_fkts.wav 10323606 bytes at offset 909098460
file #129(idx:198)/203: data/audio/music/tdr_mus_to_be_feared.wav 7999326 bytes at offset 919422068
file #130(idx:91)/203: data/audio/music/rar_mus_smash.wav 13578466 bytes at offset 927421396
file #131(idx:11)/203: data/audio/music/rab_mus_militant_force_fkts.wav 6463526 bytes at offset 940999864
file #132(idx:140)/203: data/audio/music/tdc_mus_iron_fist.wav 2522456 bytes at offset 947463392
file #133(idx:187)/203: data/audio/music/tdr_mus_no_mercy.wav 9999506 bytes at offset 949985848
file #134(idx:31)/203: data/audio/music/rac_mus_hell_march.wav 4600056 bytes at offset 959985356
file #135(idx:49)/203: data/audio/music/rac_mus_the_search.wav 3342996 bytes at offset 964585412
file #136(idx:26)/203: data/audio/music/rac_mus_face_the_enemy_2.wav 4042996 bytes at offset 967928408
file #137(idx:15)/203: data/audio/music/rab_mus_workmen_fkts.wav 12735246 bytes at offset 971971404
file #138(idx:152)/203: data/audio/music/tdc_mus_ride_of_the_valkyries.wav 3639096 bytes at offset 984706652
file #139(idx:8)/203: data/audio/music/rab_mus_hell_march_original_demo.wav 10521006 bytes at offset 988345748
file #140(idx:69)/203: data/audio/music/rar_mus_dense.wav 14597106 bytes at offset 998866756
file #141(idx:106)/203: data/audio/music/tdb_mus_act_on_instinct_fkts.wav 8126026 bytes at offset 1013463864
file #142(idx:109)/203: data/audio/music/tdb_mus_dusk_hour_fkts.wav 11315226 bytes at offset 1021589892
file #143(idx:39)/203: data/audio/music/rac_mus_radio_2.wav 2924536 bytes at offset 1032905120
file #144(idx:62)/203: data/audio/music/rar_mus_awaiting.wav 12787466 bytes at offset 1035829656
file #145(idx:70)/203: data/audio/music/rar_mus_face_the_enemy_1.wav 13387506 bytes at offset 1048617124
file #146(idx:150)/203: data/audio/music/tdc_mus_reaching_out.wav 1435426 bytes at offset 1062004632
file #147(idx:57)/203: data/audio/music/rac_mus_voice_rhythm_2.wav 3145176 bytes at offset 1063440060
file #148(idx:60)/203: data/audio/music/rac_mus_workmen_retaliation_remix.wav 2478286 bytes at offset 1066585236
file #149(idx:123)/203: data/audio/music/tdc_mus_canyon_chase.wav 459346 bytes at offset 1069063524
file #150(idx:172)/203: data/audio/music/tdr_mus_fight_win_prevail.wav 2557806 bytes at offset 1069522872
file #151(idx:107)/203: data/audio/music/tdb_mus_cc_credits_outtakes.wav 8961406 bytes at offset 1072080680
file #152(idx:183)/203: data/audio/music/tdr_mus_just_do_it_up_ost_version.wav 6894446 bytes at offset 1081042088
file #153(idx:51)/203: data/audio/music/rac_mus_traction.wav 2873366 bytes at offset 1087936536
file #154(idx:47)/203: data/audio/music/rac_mus_snake.wav 3414606 bytes at offset 1090809904
file #155(idx:23)/203: data/audio/music/rac_mus_crush_retaliation_remix.wav 2684576 bytes at offset 1094224512
file #156(idx:22)/203: data/audio/music/rac_mus_crush.wav 2739876 bytes at offset 1096909088
file #157(idx:10)/203: data/audio/music/rab_mus_journey_cs.wav 13055006 bytes at offset 1099648964
file #158(idx:115)/203: data/audio/music/tdb_mus_prepare_for_battle_fkts.wav 9816806 bytes at offset 1112703972
file #159(idx:200)/203: data/audio/music/tdr_mus_voice_rhythm.wav 14847286 bytes at offset 1122520780
file #160(idx:61)/203: data/audio/music/rar_mus_arazoid.wav 12876086 bytes at offset 1137368068
file #161(idx:93)/203: data/audio/music/rar_mus_terminate.wav 15462866 bytes at offset 1150244156
file #162(idx:79)/203: data/audio/music/rar_mus_map_theme.wav 3149586 bytes at offset 1165707024
file #163(idx:89)/203: data/audio/music/rar_mus_running_through_pipes.wav 13561666 bytes at offset 1168856612
file #164(idx:78)/203: data/audio/music/rar_mus_intro_menu.wav 10129426 bytes at offset 1182418280
file #165(idx:95)/203: data/audio/music/rar_mus_the_second_hand.wav 13301266 bytes at offset 1192547708
file #166(idx:134)/203: data/audio/music/tdc_mus_great_shot.wav 483146 bytes at offset 1205848976
file #167(idx:160)/203: data/audio/music/tdr_mus_act_on_instinct.wav 8041186 bytes at offset 1206332124
file #168(idx:137)/203: data/audio/music/tdc_mus_industrial.wav 2083626 bytes at offset 1214373312
file #169(idx:3)/203: data/audio/music/rab_mus_brain_freeze_fkts.wav 12402746 bytes at offset 1216456940
file #170(idx:83)/203: data/audio/music/rar_mus_no_mercy_retaliation_remix.wav 10252906 bytes at offset 1228859688
file #171(idx:130)/203: data/audio/music/tdc_mus_drill.wav 3237506 bytes at offset 1239112596
file #172(idx:114)/203: data/audio/music/tdb_mus_no_mercy_fkts.wav 9630746 bytes at offset 1242350104
file #173(idx:176)/203: data/audio/music/tdr_mus_heartbreak.wav 9786006 bytes at offset 1251980852
file #174(idx:21)/203: data/audio/music/rac_mus_chaos.wav 3080216 bytes at offset 1261766860
file #175(idx:104)/203: data/audio/music/rar_mus_workmen.wav 13712866 bytes at offset 1264847076
file #176(idx:52)/203: data/audio/music/rac_mus_trenches.wav 3847626 bytes at offset 1278559944
file #177(idx:12)/203: data/audio/music/rab_mus_surf_no_mercy.wav 9275986 bytes at offset 1282407572
file #178(idx:194)/203: data/audio/music/tdr_mus_recon.wav 12678126 bytes at offset 1291683560
file #179(idx:167)/203: data/audio/music/tdr_mus_demolition.wav 8823226 bytes at offset 1304361688
file #180(idx:189)/203: data/audio/music/tdr_mus_on_the_prowl.wav 8789766 bytes at offset 1313184916
file #181(idx:28)/203: data/audio/music/rac_mus_fogger.wav 3665976 bytes at offset 1321974684
file #182(idx:185)/203: data/audio/music/tdr_mus_nod_map_theme.wav 1977786 bytes at offset 1325640660
file #183(idx:42)/203: data/audio/music/rac_mus_roll_out.wav 2802176 bytes at offset 1327618448
file #184(idx:72)/203: data/audio/music/rar_mus_floating.wav 13566286 bytes at offset 1330420624
file #185(idx:129)/203: data/audio/music/tdc_mus_die.wav 1927806 bytes at offset 1343986912
file #186(idx:126)/203: data/audio/music/tdc_mus_creeping_upon.wav 2641666 bytes at offset 1345914720
file #187(idx:65)/203: data/audio/music/rar_mus_bog.wav 10454786 bytes at offset 1348556388
file #188(idx:184)/203: data/audio/music/tdr_mus_march_to_doom.wav 7580026 bytes at offset 1359011176
file #189(idx:40)/203: data/audio/music/rac_mus_radio_2_retaliation_remix.wav 2603026 bytes at offset 1366591204
file #190(idx:90)/203: data/audio/music/rar_mus_shut_it.wav 12922426 bytes at offset 1369194232
file #191(idx:124)/203: data/audio/music/tdc_mus_cc_80s_mix.wav 2955896 bytes at offset 1382116660
file #192(idx:197)/203: data/audio/music/tdr_mus_target_mechanical_man.wav 8333786 bytes at offset 1385072556
file #193(idx:81)/203: data/audio/music/rar_mus_mud.wav 13932806 bytes at offset 1393406344
file #194(idx:144)/203: data/audio/music/tdc_mus_nod_score.wav 401176 bytes at offset 1407339152
file #195(idx:132)/203: data/audio/music/tdc_mus_fight_win_prevail.wav 634556 bytes at offset 1407740328
file #196(idx:25)/203: data/audio/music/rac_mus_face_the_enemy_1.wav 3341946 bytes at offset 1408374884
file #197(idx:196)/203: data/audio/music/tdr_mus_take_em_out.wav 8785286 bytes at offset 1411716832
file #198(idx:122)/203: data/audio/music/tdc_mus_airstrike.wav 2380426 bytes at offset 1420502120
file #199(idx:131)/203: data/audio/music/tdc_mus_drone.wav 3279016 bytes at offset 1422882548
file #200(idx:71)/203: data/audio/music/rar_mus_face_the_enemy_2.wav 16221106 bytes at offset 1426161564
file #201(idx:135)/203: data/audio/music/tdc_mus_great_shot_extended.wav 678166 bytes at offset 1442382672
file #202(idx:41)/203: data/audio/music/rac_mus_reload_fire.wav 1970506 bytes at offset 1443060840
file #203(idx:37)/203: data/audio/music/rac_mus_mud_retaliation_remix.wav 3463536 bytes at offset 1445031348
```

Convert all to FLAC:

    % find . -iname "*.wav" -exec ffmpeg -hide_banner -i "{}" "{}.flac" \;
