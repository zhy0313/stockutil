package main

var testseries = [][]float64{
	{1450743840, 439.9},
	{1450743900, 439.08},
	{1450743960, 439.09},
	{1450744020, 439.09},
	{1450744080, 439.75},
	{1450744140, 439.74},
	{1450744200, 439.15},
	{1450744260, 439.89},
	{1450744320, 438.43},
	{1450744380, 439.42},
	{1450744440, 439.2},
	{1450744500, 439.5},
	{1450744680, 439.47},
	{1450744800, 439.47},
	{1450744860, 439.47},
	{1450745040, 439.49},
	{1450745100, 439.47},
	{1450745160, 439.3},
	{1450745220, 438.67},
	{1450745280, 437.55},
	{1450745340, 438.05},
	{1450745400, 438.62},
	{1450745460, 438.5},
	{1450745520, 438.5},
	{1450745580, 438.5},
	{1450745640, 438.93},
	{1450745700, 438.3},
	{1450745760, 438.3},
	{1450745880, 438.3},
	{1450745940, 438},
	{1450746000, 438.48},
	{1450746060, 438.67},
	{1450746180, 438.35},
	{1450746240, 438.35},
	{1450746300, 438.35},
	{1450746360, 439.14},
	{1450746420, 438.92},
	{1450746480, 438.9},
	{1450746540, 439.65},
	{1450746600, 439.08},
	{1450746720, 439.07},
	{1450746780, 438.91},
	{1450746840, 439.75},
	{1450747020, 439.62},
	{1450747200, 439.82},
	{1450747260, 439.83},
	{1450747380, 439.81},
	{1450747440, 439.81},
	{1450747560, 439.81},
	{1450747620, 439.81},
	{1450747680, 439.81},
	{1450747740, 439.11},
	{1450747800, 439.75},
	{1450747860, 440},
	{1450747980, 440},
	{1450748040, 439.75},
	{1450748100, 439.76},
	{1450748220, 439.76},
	{1450748280, 439.74},
	{1450748340, 439.28},
	{1450748400, 439.28},
	{1450748460, 439.13},
	{1450748520, 439.73},
	{1450748580, 439.73},
	{1450748700, 439.76},
	{1450748760, 439.76},
	{1450748820, 439.99},
	{1450748880, 439.44},
	{1450748940, 439.99},
	{1450749000, 439.45},
	{1450749060, 439.99},
	{1450749120, 439.33},
	{1450749240, 439.99},
	{1450749300, 439.99},
	{1450749360, 439.99},
	{1450749420, 439.99},
	{1450749480, 439.99},
	{1450749660, 439.99},
	{1450749720, 439.13},
	{1450749840, 439.45},
	{1450749900, 439.24},
	{1450750020, 439.99},
	{1450750080, 439.14},
	{1450750200, 439.8},
	{1450750260, 438.89},
	{1450750320, 438.44},
	{1450750380, 438.72},
	{1450750560, 438.94},
	{1450750620, 438.94},
	{1450750680, 438.19},
	{1450750740, 437.95},
	{1450750800, 438.64},
	{1450750920, 438.57},
	{1450750980, 437.91},
	{1450751040, 438.82},
	{1450751100, 439.04},
	{1450751220, 438.7},
	{1450751280, 438.97},
	{1450751340, 438.55},
	{1450751400, 439.05},
	{1450751460, 437.97},
	{1450751520, 438.43},
	{1450751580, 438.44},
	{1450751640, 437.96},
	{1450751700, 438.61},
	{1450751760, 438.61},
	{1450751880, 439.3},
	{1450751940, 438.73},
	{1450752000, 438.59},
	{1450752060, 438.73},
	{1450752120, 439.55},
	{1450752180, 439.3},
	{1450752240, 438.88},
	{1450752300, 438.67},
	{1450752360, 438.67},
	{1450752420, 437.98},
	{1450752480, 438.47},
	{1450752540, 438.5},
	{1450752600, 437.96},
	{1450752660, 438.33},
	{1450752720, 437.96},
	{1450752780, 437.96},
	{1450752840, 436.99},
	{1450752900, 437.15},
	{1450752960, 437.4},
	{1450753020, 437.32},
	{1450753080, 437.85},
	{1450753140, 437.58},
	{1450753200, 437.94},
	{1450753260, 438.51},
	{1450753320, 438.09},
	{1450753380, 438.27},
	{1450753440, 438.03},
	{1450753500, 438.13},
	{1450753560, 438.13},
	{1450753620, 438.92},
	{1450753800, 438.91},
	{1450753860, 438.91},
	{1450753920, 438.58},
	{1450753980, 438.91},
	{1450754160, 438.58},
	{1450754280, 438.82},
	{1450754340, 438.91},
	{1450754400, 438.91},
	{1450754460, 438.91},
	{1450754520, 438.99},
	{1450755060, 438.92},
	{1450755120, 438.92},
	{1450755180, 438.92},
	{1450755240, 438.58},
	{1450755300, 438.1},
	{1450755420, 438.32},
	{1450755480, 438.79},
	{1450755540, 439.33},
	{1450755600, 438.79},
	{1450755720, 439.33},
	{1450755780, 439.1},
	{1450755840, 438.95},
	{1450756020, 439.54},
	{1450756080, 439.59},
	{1450756140, 439.59},
	{1450756260, 439.59},
	{1450756320, 439.59},
	{1450756380, 439.33},
	{1450756440, 439.12},
	{1450756500, 438.42},
	{1450756560, 438.68},
	{1450756620, 438.68},
	{1450756740, 438.11},
	{1450756800, 438.2},
	{1450756860, 438.47},
	{1450756920, 438.07},
	{1450756980, 438.36},
	{1450757040, 437.9},
	{1450757100, 438.47},
	{1450757160, 438.01},
	{1450757280, 438.05},
	{1450757340, 438.47},
	{1450757460, 438.47},
	{1450757520, 439.37},
	{1450757580, 438.86},
	{1450757700, 438.86},
	{1450757760, 439.35},
	{1450757820, 438.89},
	{1450757880, 439.6},
	{1450758060, 438.86},
	{1450758120, 439.58},
	{1450758180, 439.82},
	{1450758240, 439.02},
	{1450758360, 439},
	{1450758540, 439.46},
	{1450758600, 439.15},
	{1450758660, 439.86},
	{1450758780, 439.55},
	{1450758840, 438.88},
	{1450758900, 439.46},
	{1450758960, 439.6},
	{1450759020, 439.98},
	{1450759080, 439.66},
	{1450759140, 438.87},
	{1450759200, 439.49},
	{1450759260, 438.91},
	{1450759380, 439.82},
	{1450759440, 438.97},
	{1450759500, 439.97},
	{1450759560, 439.13},
	{1450759620, 439.97},
	{1450759680, 439.91},
	{1450759800, 439.87},
	{1450760040, 439.85},
	{1450760220, 439.48},
	{1450760280, 439.2},
	{1450760340, 439.86},
	{1450760400, 439.85},
	{1450760460, 439.42},
	{1450760700, 439.71},
	{1450760760, 439.71},
	{1450760820, 438.87},
	{1450760880, 437.87},
	{1450760940, 438.76},
	{1450761060, 438.4},
	{1450761120, 438.66},
	{1450761240, 438.85},
	{1450761300, 438.38},
	{1450761360, 438.78},
	{1450761480, 438.03},
	{1450761600, 438.61},
	{1450761720, 438.2},
	{1450761900, 438.38},
	{1450761960, 438.42},
	{1450762020, 438.62},
	{1450762140, 438.61},
	{1450762200, 438.38},
	{1450762260, 437.41},
	{1450762320, 437.41},
	{1450762440, 437.41},
	{1450762560, 437.41},
	{1450762620, 437.14},
	{1450762740, 437.41},
	{1450762800, 437.41},
	{1450762860, 437.41},
	{1450762920, 437.41},
	{1450762980, 437.41},
	{1450763040, 437.16},
	{1450763100, 437.41},
	{1450763160, 437.16},
	{1450763220, 437.41},
	{1450763280, 437.41},
	{1450763340, 437.19},
	{1450763460, 438.11},
	{1450763580, 437.74},
	{1450763640, 438.02},
	{1450763760, 437.99},
	{1450763820, 437.99},
	{1450763940, 438.43},
	{1450764000, 438.41},
	{1450764060, 438.41},
	{1450764300, 438.41},
	{1450764360, 438.42},
	{1450764420, 438.24},
	{1450764480, 438},
	{1450764540, 437.98},
	{1450764720, 437.34},
	{1450764840, 438.55},
	{1450764900, 437.59},
	{1450765020, 438.68},
	{1450765380, 438.61},
	{1450765440, 439.04},
	{1450765560, 439.02},
	{1450765680, 439.03},
	{1450765800, 439.07},
	{1450765860, 439.08},
	{1450766100, 439.06},
	{1450766220, 438.24},
	{1450766280, 438.24},
	{1450766340, 438.24},
	{1450766400, 437.9},
	{1450766520, 438.24},
	{1450766700, 438.43},
	{1450766940, 439.22},
	{1450767120, 439},
	{1450767240, 438.36},
	{1450767300, 437.98},
	{1450767420, 437.99},
	{1450767480, 437.9},
	{1450767540, 437.53},
	{1450767600, 438.07},
	{1450767780, 437.96},
	{1450767900, 437.96},
	{1450767960, 439.11},
	{1450768080, 439.02},
	{1450768320, 438.05},
	{1450768440, 438},
	{1450768500, 438},
	{1450768680, 438},
	{1450768740, 438},
	{1450768800, 438},
	{1450768860, 438.42},
	{1450768920, 438.88},
	{1450769160, 439},
	{1450769220, 439.07},
	{1450769340, 439.68},
	{1450769400, 439.99},
	{1450769460, 441},
	{1450769520, 441},
	{1450769580, 440.02},
	{1450769640, 440.01},
	{1450769760, 440.01},
	{1450769880, 440.6},
	{1450770000, 441.02},
	{1450770060, 441.5},
	{1450770240, 441.26},
	{1450770300, 440.99},
	{1450770360, 441.12},
	{1450770420, 441.53},
	{1450770480, 441.5},
	{1450770540, 440.8},
	{1450770600, 440.54},
	{1450770660, 440.51},
	{1450770900, 440.99},
	{1450770960, 440.99},
	{1450771080, 440.99},
	{1450771200, 441},
	{1450771320, 441.23},
	{1450771440, 441.02},
	{1450771560, 440.51},
	{1450771620, 440.86},
	{1450771740, 440.49},
	{1450771800, 440.49},
	{1450771920, 440.85},
	{1450771980, 440.77},
	{1450772040, 441.03},
	{1450772100, 440.85},
	{1450772160, 440.97},
	{1450772220, 440.85},
	{1450772280, 441.03},
	{1450772340, 441.47},
	{1450772400, 440.74},
	{1450772460, 441.17},
	{1450772520, 441.48},
	{1450772580, 441.53},
	{1450772640, 441.5},
	{1450772700, 440.76},
	{1450772820, 441.21},
	{1450772940, 441},
	{1450773000, 441.52},
	{1450773060, 441.52},
	{1450773120, 441.79},
	{1450773180, 441.8},
	{1450773240, 441.98},
	{1450773300, 441.82},
	{1450773360, 441.82},
	{1450773420, 442.87},
	{1450773660, 442.66},
	{1450773720, 442.65},
	{1450773780, 442.25},
	{1450773960, 442.24},
	{1450774020, 442.24},
	{1450774140, 442.24},
	{1450774200, 442.17},
	{1450774560, 441.98},
	{1450774680, 441.99},
	{1450774800, 441.79},
	{1450774860, 441.79},
	{1450774920, 441.79},
	{1450774980, 442},
	{1450775040, 442.43},
	{1450775100, 442.39},
	{1450775220, 442.39},
	{1450775280, 442.39},
	{1450775340, 442},
	{1450775400, 442.59},
	{1450775460, 442.39},
	{1450775520, 442},
	{1450775580, 442},
	{1450775700, 442.47},
	{1450775820, 441.5},
	{1450775880, 442.35},
	{1450775940, 442.64},
	{1450776060, 442.37},
	{1450776120, 442.37},
	{1450776180, 442.36},
	{1450776240, 442.37},
	{1450776360, 441.99},
	{1450776420, 441.99},
	{1450776600, 441.99},
	{1450776660, 441.99},
	{1450776720, 441.99},
	{1450776780, 441.89},
	{1450776840, 441.99},
	{1450776900, 442.38},
	{1450777140, 442.22},
	{1450777200, 441.83},
	{1450777260, 441.82},
	{1450777380, 442.15},
	{1450777440, 442.39},
	{1450777500, 442.39},
	{1450777740, 441.82},
	{1450777860, 442.19},
	{1450777920, 442.19},
	{1450778040, 442.43},
	{1450778100, 443.49},
	{1450778160, 442.11},
	{1450778220, 443.18},
	{1450778280, 443.4},
	{1450778400, 442.79},
	{1450778520, 442.16},
	{1450778580, 442.39},
	{1450778760, 442.27},
	{1450778820, 442.51},
	{1450778880, 442.75},
	{1450778940, 443.38},
	{1450779000, 443.39},
	{1450779120, 443.4},
	{1450779180, 443.4},
	{1450779240, 443},
	{1450779300, 442.8},
	{1450779360, 442.5},
	{1450779420, 441.81},
	{1450779480, 441.81},
	{1450779540, 441.76},
	{1450779600, 441.53},
	{1450779660, 441.46},
	{1450779780, 441.64},
	{1450780020, 441.71},
	{1450780140, 441.01},
	{1450780200, 441.69},
	{1450780380, 441.71},
	{1450780560, 441.71},
	{1450780620, 441.05},
	{1450780680, 441.01},
	{1450780740, 441.7},
	{1450780800, 441},
	{1450780980, 441.27},
	{1450781040, 441.26},
	{1450781100, 441},
	{1450781160, 441},
	{1450781220, 441},
	{1450781280, 441.2},
	{1450781400, 441.2},
	{1450781460, 440.79},
	{1450781520, 440.27},
	{1450781580, 440.27},
	{1450781640, 440.27},
	{1450781700, 440.32},
	{1450781760, 440.77},
	{1450781940, 441.08},
	{1450782000, 441},
	{1450782060, 440.27},
	{1450782120, 439.75},
	{1450782180, 439.6},
	{1450782240, 439.53},
	{1450782300, 439.57},
	{1450782360, 439.51},
	{1450782420, 439.3},
	{1450782480, 439.89},
	{1450782540, 439.89},
	{1450782600, 440.9},
	{1450782660, 440.83},
	{1450782720, 440.03},
	{1450782780, 440.02},
	{1450782840, 440.48},
	{1450782960, 441.43},
	{1450783020, 441.44},
	{1450783140, 441.2},
	{1450783200, 441},
	{1450783260, 440.78},
	{1450783380, 440.77},
	{1450783500, 441.15},
	{1450783560, 441.16},
	{1450783800, 440.48},
	{1450783980, 440.86},
	{1450784040, 440.86},
	{1450784220, 440.58},
	{1450784280, 440.37},
	{1450784340, 440.92},
	{1450784400, 440.93},
	{1450784460, 440.85},
	{1450784700, 440.9},
	{1450784760, 440.9},
	{1450784820, 440.98},
	{1450784880, 440.8},
	{1450784940, 440.28},
	{1450785000, 440.29},
	{1450785060, 440.2},
	{1450785120, 440.98},
	{1450785180, 440.99},
	{1450785240, 440.99},
	{1450785300, 441},
	{1450785360, 440.99},
	{1450785420, 441},
	{1450785480, 441.42},
	{1450785540, 441.4},
	{1450785600, 440.57},
	{1450785660, 441.01},
	{1450785720, 441.01},
	{1450785780, 440.94},
	{1450785840, 441.42},
	{1450785900, 441.42},
	{1450785960, 441.42},
	{1450786020, 440.94},
	{1450786140, 440.94},
	{1450786200, 440.94},
	{1450786260, 440.94},
	{1450786320, 440.94},
	{1450786380, 441.28},
	{1450786440, 441},
	{1450786500, 441.02},
	{1450786560, 441.02},
	{1450786620, 440.98},
	{1450786680, 440.96},
	{1450786740, 440.6},
	{1450786800, 440.51},
	{1450786860, 440.51},
	{1450786920, 440.51},
	{1450786980, 440.51},
	{1450787040, 440.51},
	{1450787220, 440.99},
	{1450787340, 441},
	{1450787400, 441},
	{1450787580, 441},
	{1450787640, 440.58},
	{1450787700, 440.58},
	{1450787760, 440.5},
	{1450787820, 441},
	{1450787940, 441},
	{1450788060, 440.5},
	{1450788180, 440.58},
	{1450788240, 440.92},
	{1450788300, 441},
	{1450788360, 440.5},
	{1450788420, 441},
	{1450788480, 441},
	{1450788540, 441},
	{1450788660, 440.5},
	{1450788720, 441},
	{1450788780, 440.9},
	{1450788840, 441},
	{1450789020, 441},
	{1450789260, 441},
	{1450789500, 441},
	{1450789560, 441},
	{1450789620, 440.6},
	{1450789680, 440.6},
	{1450789800, 441},
	{1450789860, 441},
	{1450790040, 440.6},
	{1450790100, 440.84},
	{1450790220, 440.95},
	{1450790280, 440.84},
	{1450790340, 441},
	{1450790400, 440.11},
	{1450790460, 440},
	{1450790520, 439.01},
	{1450790580, 438.92},
	{1450790640, 438.8},
	{1450790760, 438.51},
	{1450790820, 436.29},
	{1450790880, 437.6},
	{1450790940, 437.54},
	{1450791000, 438},
	{1450791060, 438.53},
	{1450791120, 438.47},
	{1450791180, 438.59},
	{1450791240, 438.6},
	{1450791300, 438.59},
	{1450791360, 439},
	{1450791420, 439.01},
	{1450791480, 438.67},
	{1450791540, 438.61},
	{1450791600, 437.99},
	{1450791660, 438.51},
	{1450791720, 439.1},
	{1450791780, 439.19},
	{1450791840, 438.9},
	{1450791960, 438.59},
	{1450792020, 438.5},
	{1450792080, 438.5},
	{1450792140, 438.5},
	{1450792200, 438.82},
	{1450792260, 437.39},
	{1450792320, 437.27},
	{1450792380, 437.1},
	{1450792440, 437.07},
	{1450792500, 436.99},
	{1450792560, 436.86},
	{1450792620, 436.76},
	{1450792680, 436.33},
	{1450792740, 436.75},
	{1450792800, 436.63},
	{1450792860, 436.74},
	{1450792920, 436.5},
	{1450792980, 434.76},
	{1450793040, 434.77},
	{1450793100, 435.3},
	{1450793160, 435.59},
	{1450793220, 434.67},
	{1450793280, 435.6},
	{1450793340, 436.55},
	{1450793400, 435.65},
	{1450793460, 435.59},
	{1450793520, 435.63},
	{1450793580, 436.63},
	{1450793640, 436.76},
	{1450793760, 436.14},
	{1450793820, 436.15},
	{1450793880, 436.76},
	{1450793940, 437.1},
	{1450794000, 436.45},
	{1450794060, 436.6},
	{1450794120, 436.42},
	{1450794180, 436.22},
	{1450794240, 436.41},
	{1450794300, 436.43},
	{1450794360, 436.41},
	{1450794420, 436.45},
	{1450794480, 436.45},
	{1450794540, 436.44},
	{1450794600, 436.4},
	{1450794660, 436.45},
	{1450794720, 436.45},
	{1450794780, 436.45},
	{1450794840, 436.45},
	{1450794900, 436.45},
	{1450794960, 436.45},
	{1450795020, 436.45},
	{1450795080, 436.45},
	{1450795140, 435.99},
	{1450795200, 436.46},
	{1450795260, 436.27},
	{1450795320, 436.57},
	{1450795380, 436.81},
	{1450795440, 436.27},
	{1450795500, 436.58},
	{1450795560, 436.27},
	{1450795620, 435.84},
	{1450795680, 435.83},
	{1450795740, 435.87},
	{1450795800, 435.69},
	{1450795860, 435.91},
	{1450795980, 435.94},
	{1450796100, 436.4},
	{1450796160, 436.18},
	{1450796220, 436.67},
	{1450796280, 437.08},
	{1450796340, 437.08},
	{1450796400, 436.85},
	{1450796460, 436.22},
	{1450796520, 435.46},
	{1450796580, 435.46},
	{1450796640, 435.46},
	{1450796700, 435.46},
	{1450796760, 435.46},
	{1450796820, 435.46},
	{1450796880, 435.46},
	{1450796940, 435.37},
	{1450797060, 435.46},
	{1450797120, 435.46},
	{1450797180, 435.46},
	{1450797240, 435.46},
	{1450797300, 435.46},
	{1450797360, 435.46},
	{1450797420, 435.46},
	{1450797480, 435.46},
	{1450797540, 435.46},
	{1450797600, 435.43},
	{1450797660, 435.19},
	{1450797720, 434.88},
	{1450797780, 435.19},
	{1450797840, 434.91},
	{1450797900, 434.91},
	{1450797960, 434.78},
	{1450798020, 435.19},
	{1450798080, 435.19},
	{1450798140, 435.19},
	{1450798200, 435.18},
	{1450798260, 435.19},
	{1450798320, 435.19},
	{1450798380, 435.19},
	{1450798440, 435.19},
	{1450798500, 435.19},
	{1450798560, 435.19},
	{1450798620, 435.25},
	{1450798680, 435.81},
	{1450798740, 435.79},
	{1450798860, 435.91},
	{1450798920, 435.84},
	{1450798980, 435.97},
	{1450799100, 435.99},
	{1450799160, 436.27},
	{1450799220, 436.28},
	{1450799280, 436.89},
	{1450799340, 436.62},
	{1450799400, 437.13},
	{1450799460, 436.64},
	{1450799520, 436.84},
	{1450799580, 437.11},
	{1450799700, 437.18},
	{1450799760, 436.64},
	{1450799820, 436.87},
	{1450799880, 436.93},
	{1450799940, 437.2},
	{1450800000, 436.87},
	{1450800060, 436.99},
	{1450800120, 436.88},
	{1450800180, 436.89},
	{1450800240, 437.49},
	{1450800300, 437.11},
	{1450800360, 438.05},
	{1450800420, 437.42},
	{1450800480, 437.75},
	{1450800540, 437.21},
	{1450800600, 436.88},
	{1450800840, 437.15},
	{1450800900, 437.73},
	{1450800960, 437.2},
	{1450801020, 437},
	{1450801080, 437.17},
	{1450801140, 437.2},
	{1450801260, 437.2},
	{1450801320, 437.62},
	{1450801380, 437.72},
	{1450801440, 437.74},
	{1450801500, 437.74},
	{1450801560, 437.74},
	{1450801620, 437.74},
	{1450801680, 436.83},
	{1450801740, 437.69},
	{1450801800, 437.75},
	{1450801860, 437.75},
	{1450801920, 437.74},
	{1450801980, 438.09},
	{1450802040, 438.25},
	{1450802100, 438.25},
	{1450802160, 438.09},
	{1450802220, 438.25},
	{1450802280, 438.25},
	{1450802400, 438.09},
	{1450802460, 437.19},
	{1450802520, 437.64},
	{1450802580, 437.64},
	{1450802640, 437.64},
	{1450802700, 437.64},
	{1450802760, 437.64},
	{1450802820, 437.47},
	{1450802880, 437.64},
	{1450802940, 437.64},
	{1450803000, 437.64},
	{1450803060, 437.9},
	{1450803120, 438.25},
	{1450803180, 438.25},
	{1450803240, 438.25},
	{1450803360, 437.99},
	{1450803420, 437.87},
	{1450803480, 437.22},
	{1450803600, 437.81},
	{1450803660, 437.84},
	{1450803720, 437.88},
	{1450803780, 437.88},
	{1450803840, 437.57},
	{1450803900, 437.06},
	{1450803960, 437.61},
	{1450804020, 437.56},
	{1450804080, 437.59},
	{1450804200, 437.56},
	{1450804260, 437.8},
	{1450804320, 437.56},
	{1450804380, 437.74},
	{1450804440, 437.22},
	{1450804500, 437.74},
	{1450804560, 437.74},
	{1450804620, 437.88},
	{1450804680, 437.88},
	{1450804740, 437.55},
	{1450804800, 437.6},
	{1450804860, 437.36},
	{1450804920, 437.08},
	{1450804980, 436.63},
	{1450805040, 437.16},
	{1450805100, 437.16},
	{1450805160, 437.14},
	{1450805220, 437.16},
	{1450805280, 437.16},
	{1450805340, 437.16},
	{1450805400, 437.16},
	{1450805460, 436.92},
	{1450805520, 436.66},
	{1450805580, 433.81},
	{1450805640, 433.77},
	{1450805700, 434.57},
	{1450805760, 434.57},
	{1450805820, 435.09},
	{1450805880, 435.08},
	{1450805940, 435.05},
	{1450806000, 435.84},
	{1450806060, 435.45},
	{1450806120, 434.64},
	{1450806180, 434.47},
	{1450806240, 435.05},
	{1450806300, 435.83},
	{1450806360, 435.82},
	{1450806420, 435.41},
	{1450806480, 435.82},
	{1450806540, 435.72},
	{1450806600, 434.51},
	{1450806840, 435.05},
	{1450806900, 435.05},
	{1450806960, 436.12},
	{1450807020, 436.12},
	{1450807080, 435.49},
	{1450807140, 435.52},
	{1450807200, 436.12},
	{1450807260, 436.94},
	{1450807320, 436.22},
	{1450807380, 436.22},
	{1450807440, 436.22},
	{1450807500, 435.76},
	{1450807560, 435.87},
	{1450807680, 435.98},
	{1450807740, 435.98},
	{1450807860, 436.05},
	{1450808100, 436},
	{1450808160, 436},
	{1450808220, 436.25},
	{1450808280, 436.25},
	{1450808340, 436.25},
	{1450808400, 436.25},
	{1450808460, 436.25},
	{1450808520, 435.97},
	{1450808580, 435.25},
	{1450808640, 435.57},
	{1450808760, 435.95},
	{1450808820, 435.16},
	{1450808880, 434.96},
	{1450808940, 434.64},
	{1450809000, 435.18},
	{1450809120, 435.18},
	{1450809180, 435.09},
	{1450809300, 435.17},
	{1450809360, 435.95},
	{1450809420, 435.68},
	{1450809480, 435.37},
	{1450809540, 434.64},
	{1450809600, 434.75},
	{1450809660, 434.7},
	{1450809780, 435.36},
	{1450809840, 435.37},
	{1450809900, 435.37},
	{1450809960, 435.47},
	{1450810020, 435.47},
	{1450810080, 435.5},
	{1450810140, 435},
	{1450810200, 434.76},
	{1450810260, 435.48},
	{1450810320, 435.33},
	{1450810380, 434.82},
	{1450810440, 434.81},
	{1450810500, 434.81},
	{1450810560, 434.76},
	{1450810620, 434.81},
	{1450810680, 434.85},
	{1450810740, 433.66},
	{1450810800, 433.53},
	{1450810860, 434.84},
	{1450810920, 434.85},
	{1450810980, 434.84},
	{1450811040, 434.03},
	{1450811100, 435.3},
	{1450811160, 435.47},
	{1450811220, 434.78},
	{1450811280, 434.78},
	{1450811340, 434.7},
	{1450811400, 434.81},
	{1450811460, 435.06},
	{1450811520, 434.71},
	{1450811580, 435.06},
	{1450811640, 435.27},
	{1450811700, 435.25},
	{1450811760, 434.71},
	{1450811820, 434.67},
	{1450811880, 434.62},
	{1450811940, 434.68},
	{1450812060, 434.69},
	{1450812120, 434.69},
	{1450812180, 434.69},
	{1450812240, 434.69},
	{1450812300, 434.69},
	{1450812360, 435.16},
	{1450812420, 435.38},
	{1450812480, 435.25},
	{1450812540, 435.19},
	{1450812660, 435.23},
	{1450812720, 435.22},
	{1450812780, 435.76},
	{1450812840, 435.71},
	{1450812960, 435.7},
	{1450813020, 435.67},
	{1450813080, 435.76},
	{1450813140, 435.73},
	{1450813260, 435.69}}