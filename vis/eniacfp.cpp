#include <irrlicht.h>
#include <iostream>
#include <thread>
#include <string>
#include <math.h>

using namespace irr;
using namespace core;
using namespace scene;
using namespace video;
using namespace io;
using namespace gui;

ICursorControl *curs;
IVolumeLightSceneNode *adneon[20][11];
IVolumeLightSceneNode *acneon[20][10];
IVolumeLightSceneNode *cycneon, *cycneon2;
IVolumeLightSceneNode *mpdneon[20], *mpsneon[10];
IVolumeLightSceneNode *fttenneon[3], *ftoneneon[3], *ftsetneon[3];
IVolumeLightSceneNode *ftaddneon[3], *ftsubneon[3], *ftringneon[3];
IVolumeLightSceneNode *consneon[20];
IVolumeLightSceneNode *mulsneon, *mulr1neon, *mulr3neon;
IVolumeLightSceneNode *dsqplneon, *dsqstat[30];
ICameraSceneNode *camera;
int laccmpos[9] = { 7725, 8335, 9555, 10165, 10775, 11385, 11995, 12605, 13215 };
int baccmpos[5] = { -2230, 210, 820, 1430, 2040 };
int raccmpos[6] = { 13690, 13080, 12470, 11860, 8810, 8200 };

void
setcam(vector3df pos, double angle) {
	camera->setPosition(vector3df(pos.X, 0, pos.Z));
	camera->updateAbsolutePosition();
}


class MyReceiver : public IEventReceiver {
public:
	MyReceiver() {}
	virtual bool OnEvent(const SEvent& event) {
		static vector3df campos(0,0,1000);
		static double angle = 0; 
		static vector3df forward(0,0,1);
		static vector3df target(0,0,20000);

		if(event.EventType == EET_MOUSE_INPUT_EVENT) {
			if(event.MouseInput.isRightPressed() && event.MouseInput.Shift) {
				curs->setPosition(-20, 0);
				return true;
			}
		}
		else if(event.EventType == EET_KEY_INPUT_EVENT) {
			if(!event.KeyInput.PressedDown)
				return false;
			else if(event.KeyInput.Key == KEY_UP) {
				campos += forward * 40;
				setcam(campos, angle);
			}
			else if(event.KeyInput.Key == KEY_DOWN) {
				campos -= forward * 40;
				setcam(campos, angle);
			}
			else if(event.KeyInput.Key == KEY_LEFT) {
				angle -= 0.015;
				target = vector3df(20000.0 * sin(angle), 0.0, 20000.0 * cos(angle));
				forward = (target - campos).normalize();
				camera->setTarget(target);
				setcam(campos, angle);
			}
			else if(event.KeyInput.Key == KEY_RIGHT) {
				angle += 0.015;
				target = vector3df(20000.0 * sin(angle), 0.0, 20000.0 * cos(angle));
				forward = (target - campos).normalize();
				camera->setTarget(target);
				setcam(campos, angle);
			}
			else if(event.KeyInput.Char == 'Q')
				exit(0);
		}
		return false;
	}
};

void
stdinreader() {
	std::string msg;
	int unit, digit, val;
	int ms, mr1, mr3;
	float xpos, ystart, dir;
	char dstat[32];

	while(1) {
		std::getline(std::cin, msg);
		if(sscanf(msg.c_str(), "ad %d %d %d", &unit, &digit, &val) == 3) {
			if(unit < 9)
				adneon[unit][digit]->setPosition(vector3df(-2690, 390 + val * 35, laccmpos[unit] + 47 * digit));
			else if(unit < 14)
				adneon[unit][digit]->setPosition(vector3df(baccmpos[unit-9] + 47 * digit, 390 + val * 35, 14150));
			else
				adneon[unit][digit]->setPosition(vector3df(2972, 390 + val * 35, raccmpos[unit-14] - 47 * digit));
		}
		else if(sscanf(msg.c_str(), "ac %d %d %d", &unit, &digit, &val) == 3) {
			if(unit < 9)
				acneon[unit][digit]->setPosition(vector3df(-2690, -200 + val * 427, laccmpos[unit] + 47 * digit));
			else if(unit < 14)
				acneon[unit][digit]->setPosition(vector3df(baccmpos[unit-9] + 47 * digit, -200 + val * 427, 14150));
			else
				acneon[unit][digit]->setPosition(vector3df(2972, -200 + val * 427, raccmpos[unit-14] - 47 * digit));
		}
		else if(sscanf(msg.c_str(), "cy %d", &val) == 1) {
			val &= ~1;
			cycneon->setPosition(vector3df(-2645, 300, 4732 + 9.6 * val));
			if(val <= 20)
				cycneon2->setPosition(vector3df(-2645, 245, 4866));
			else if(val <= 36)
				cycneon2->setPosition(vector3df(-2645, 245, 4965));
			else
				cycneon2->setPosition(vector3df(-2800, 245, 4866));
		}
		else if(sscanf(msg.c_str(), "mpd %d %d", &digit, &val) == 2) {
			if(digit < 10) {
				mpdneon[digit]->setPosition(vector3df(-2745, 475 + 20 * val, 5368 + 40 * digit));
			}
			else {
				mpdneon[digit]->setPosition(vector3df(-2745, 475 + 20 * val, 5970 + 40 * (digit - 10)));
			}
		}
		else if(sscanf(msg.c_str(), "mps %d %d", &digit, &val) == 2) {
			if(digit < 5) {
				mpsneon[digit]->setPosition(vector3df(-2745, 80 + 20 * val, 5367 + 75 * digit));
			}
			else {
				mpsneon[digit]->setPosition(vector3df(-2745, 80 + 20 * val, 5973 + 75 * (digit - 5)));
			}
		}
		else if(sscanf(msg.c_str(), "ftar %d %d", &unit, &val) == 2) {
			switch(unit) {
			case 0:
				dir = 1;
				xpos = -2645;
				ystart = 6520;
				break;
			case 1:
				dir = -1;
				xpos = 2922;
				ystart = 11230;
				break;
			case 2:
				dir = -1;
				xpos = 2922;
				ystart = 10010;
				break;
			}
			ftoneneon[unit]->setPosition(vector3df(xpos, 300, ystart + (val / 10) * 19.2 * dir));
			fttenneon[unit]->setPosition(vector3df(xpos, 300, ystart + (val % 10) * 19.2 * dir + 250 * dir));
		}
		else if(sscanf(msg.c_str(), "ftr %d %d", &unit, &val) == 2) {
			switch(unit) {
			case 0:
				dir = 1;
				xpos = -2645;
				ystart = 6737;
				break;
			case 1:
				dir = -1;
				xpos = 2922;
				ystart = 11010;
				break;
			case 2:
				dir = -1;
				xpos = 2922;
				ystart = 9790;
				break;
			}
			ftringneon[unit]->setPosition(vector3df(xpos, 245, ystart + val * 18.5 * dir));
		}
		else if(sscanf(msg.c_str(), "ftad %d %d", &unit, &val) == 2) {
			switch(unit) {
			case 0:
				xpos = -2745 + val * 100;
				ystart = 6612;
				break;
			case 1:
				xpos = 3022 - val * 100;
				ystart = 11138;
				break;
			case 2:
				xpos = 3022 - val * 100;
				ystart = 9918;
				break;
			}
			ftaddneon[unit]->setPosition(vector3df(xpos, 245, ystart));
		}
		else if(sscanf(msg.c_str(), "ftsu %d %d", &unit, &val) == 2) {
			switch(unit) {
			case 0:
				xpos = -2745 + val * 100;
				ystart = 6640;
				break;
			case 1:
				xpos = 3022 - val * 100;
				ystart = 11110;
				break;
			case 2:
				xpos = 3022 - val * 100;
				ystart = 9890;
				break;
			}
			ftsubneon[unit]->setPosition(vector3df(xpos, 245, ystart));
		}
		else if(sscanf(msg.c_str(), "ftse %d %d", &unit, &val) == 2) {
			switch(unit) {
			case 0:
				xpos = -2745 + val * 100;
				ystart = 6520;
				break;
			case 1:
				xpos = 3022 - val * 100;
				ystart = 11230;
				break;
			case 2:
				xpos = 3022 - val * 100;
				ystart = 10010;
				break;
			}
			ftsubneon[unit]->setPosition(vector3df(xpos, 245, ystart));
		}
		else if(sscanf(msg.c_str(), "ct %d %d", &digit, &val) == 2) {
			if(digit < 10) {
				consneon[digit]->setPosition(vector3df(3095 - 100 * val, 660, 7570 - 49 * (digit - 1)));
			}
			else if(digit < 20) {
				consneon[digit]->setPosition(vector3df(3095 - 100 * val, 204, 7570 - 49 * (digit - 11)));
			}
		}
		else if(sscanf(msg.c_str(), "m %d %*s %d %d", &ms, &mr1, &mr3) == 3) {
			mulsneon->setPosition(vector3df(-910 + 20 * ms, 225, 14090));
			mulr1neon->setPosition(vector3df(-1400, 230, 14190 - mr1 * 100));
			mulr3neon->setPosition(vector3df(-180, 230, 14190 - mr1 * 100));
		}
		else if(sscanf(msg.c_str(), "d %d %*d %*s %s", &val, dstat) == 2) {
			dsqplneon->setPosition(vector3df(-2720, 190 - val * 20, 8970));
			dsqstat[0]->setPosition(vector3df(-2720 + 100 * (dstat[0] - '0'), 370, 9080));
			dsqstat[1]->setPosition(vector3df(-2720 + 100 * (dstat[1] - '0'), 370, 9101));
			dsqstat[2]->setPosition(vector3df(-2720 + 100 * (dstat[2] - '0'), 370, 9122));
			dsqstat[3]->setPosition(vector3df(-2720 + 100 * (dstat[3] - '0'), 370, 9143));
			dsqstat[4]->setPosition(vector3df(-2720 + 100 * (dstat[4] - '0'), 370, 9164));
			dsqstat[5]->setPosition(vector3df(-2720 + 100 * (dstat[5] - '0'), 370, 9185));
			dsqstat[6]->setPosition(vector3df(-2720 + 100 * (dstat[6] - '0'), 370, 9206));
			dsqstat[7]->setPosition(vector3df(-2720 + 100 * (dstat[7] - '0'), 370, 9227));
			dsqstat[8]->setPosition(vector3df(-2720 + 100 * (dstat[9] - '0'), 370, 9238));
			dsqstat[9]->setPosition(vector3df(-2720 + 100 * (dstat[8] - '0'), 370, 9379));
			dsqstat[10]->setPosition(vector3df(-2820 + 100 * (dstat[10] - '0'), 190, 9010));
			dsqstat[11]->setPosition(vector3df(-2820 + 100 * (dstat[11] - '0'), 190, 9055));
			dsqstat[12]->setPosition(vector3df(-2820 + 100 * (dstat[12] - '0'), 190, 9100));
			dsqstat[13]->setPosition(vector3df(-2820 + 100 * (dstat[13] - '0'), 190, 9145));
			dsqstat[14]->setPosition(vector3df(-2820 + 100 * (dstat[14] - '0'), 190, 9190));
			dsqstat[15]->setPosition(vector3df(-2820 + 100 * (dstat[15] - '0'), 190, 9235));
			dsqstat[16]->setPosition(vector3df(-2820 + 100 * (dstat[16] - '0'), 190, 9280));
			dsqstat[17]->setPosition(vector3df(-2820 + 100 * (dstat[17] - '0'), 190, 9325));
			dsqstat[18]->setPosition(vector3df(-2820 + 100 * (dstat[27] - '0'), 190, 9370));
			dsqstat[19]->setPosition(vector3df(-2820 + 100 * (dstat[29] - '0'), 190, 9415));
			dsqstat[20]->setPosition(vector3df(-2820 + 100 * (dstat[18] - '0'), 165, 9010));
			dsqstat[21]->setPosition(vector3df(-2820 + 100 * (dstat[19] - '0'), 165, 9055));
			dsqstat[22]->setPosition(vector3df(-2820 + 100 * (dstat[20] - '0'), 165, 9100));
			dsqstat[23]->setPosition(vector3df(-2820 + 100 * (dstat[21] - '0'), 165, 9145));
			dsqstat[24]->setPosition(vector3df(-2820 + 100 * (dstat[22] - '0'), 165, 9190));
			dsqstat[25]->setPosition(vector3df(-2820 + 100 * (dstat[23] - '0'), 165, 9235));
			dsqstat[26]->setPosition(vector3df(-2820 + 100 * (dstat[24] - '0'), 165, 9280));
			dsqstat[27]->setPosition(vector3df(-2820 + 100 * (dstat[25] - '0'), 165, 9325));
			dsqstat[28]->setPosition(vector3df(-2820 + 100 * (dstat[26] - '0'), 165, 9370));
			dsqstat[29]->setPosition(vector3df(-2820 + 100 * (dstat[28] - '0'), 165, 9415));
		}
	}
}

int
main() {
	IrrlichtDevice *device;
	IVideoDriver *driver;
	ISceneManager *smgr;
	IGUIEnvironment *guienv;
	IAnimatedMesh *mesh;
	IMeshSceneNode *node;
	ITexture *texture;
	ILightSceneNode *light[4];
	MyReceiver receiver;
	int i, j;

	std::thread stdint(stdinreader);

	device = createDevice(video::EDT_OPENGL, dimension2d<u32>(1680,1050), 16,
		false, false, false, &receiver);
//	device = createDevice(video::EDT_SOFTWARE, dimension2d<u32>(1680,1050), 16,
//		false, false, false, &receiver);
	if(device == NULL) {
		perror("create device");
		exit(1);
	}
	device->setWindowCaption(L"3D ENIAC");

	driver = device->getVideoDriver();
	smgr = device->getSceneManager();
	guienv = device->getGUIEnvironment();
	curs = device->getCursorControl();
	curs->setVisible(false);

	mesh = smgr->getMesh("obj/eniact.obj");
	if(mesh == NULL) {
		perror("mesh");
		device->drop();
		exit(1);
	}
	node = smgr->addMeshSceneNode(mesh->getMesh(0));
	if(node == NULL) {
		perror("node");
		device->drop();
		exit(1);
	}
	node->setRotation(vector3df(-90, 180, 0));
	node->setPosition(vector3df(-2300, -1600, 4000));
//	node->setMaterialFlag(EMF_LIGHTING, false);
	camera = smgr->addCameraSceneNode(0, vector3df(0, 0, 1000), vector3df(0, 0, 20000));
	camera->setFOV(0.8);
	camera->bindTargetAndRotation(true);
	camera->setFarValue(20000.0);
	camera->setAspectRatio(16.0/9.0);
	light[0] = smgr->addLightSceneNode(0, vector3df(-1300, 2000, 9000), SColorf(0.6, 0.6, 0.6), 10000.0);
	light[0]->setLightType(ELT_POINT);
	light[0]->setVisible(true);
	light[1] = smgr->addLightSceneNode(0, vector3df(-1300, 2000, 3000), SColorf(0.6, 0.6, 0.6), 10000.0);
	light[1]->setLightType(ELT_POINT);
	light[1]->setVisible(true);
	light[2] = smgr->addLightSceneNode(0, vector3df(1300, 2000, 9000), SColorf(0.6, 0.6, 0.6), 10000.0);
	light[2]->setLightType(ELT_POINT);
	light[2]->setVisible(true);
	light[3] = smgr->addLightSceneNode(0, vector3df(1300, 2000, 3000), SColorf(0.6, 0.6, 0.6), 10000.0);
	light[3]->setLightType(ELT_POINT);
	light[3]->setVisible(true);

	for(i = 0; i < 9; i++) {
		for(j = 0; j < 11; j++) {
			adneon[i][j] = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
				vector3df(-2690, 390, laccmpos[i] + 47 * j), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
		}
		for(j = 1; j <= 10; j++) {
			acneon[i][j] = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
				vector3df(-2690, -200, laccmpos[i] + 47 * j), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
		}
	}
	for(i = 0; i < 5; i++) {
		for(j = 0; j < 11; j++) {
			adneon[i+9][j] = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
				vector3df(baccmpos[i] + 47 * j, 390, 14150), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
		}
		for(j = 1; j <= 10; j++) {
			acneon[i+9][j] = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
				vector3df(baccmpos[i] + 47 * j, -200, 14150), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
		}
	}
	for(i = 0; i < 6; i++) {
		for(j = 0; j < 11; j++) {
			adneon[i+14][j] = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
				vector3df(2972, 390, raccmpos[i] - 47 * j), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
		}
		for(j = 1; j <= 10; j++) {
			acneon[i+14][j] = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
				vector3df(2972, -200, raccmpos[i] - 47 * j), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
		}
	}
	cycneon = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
		vector3df(-2645, 280, 4732), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
	cycneon2 = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
		vector3df(-2645, 200, 4866), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
	for(i = 0; i < 5; i++) {
		mpsneon[i] = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
			vector3df(-2745, 80, 5367 + 75 * i), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
		mpsneon[i+5] = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
			vector3df(-2745, 80, 5973 + 75 * i), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
	}
	for(i = 0; i < 10; i++) {
		mpdneon[i] = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
			vector3df(-2745, 475, 5368 + 40 * i), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
		mpdneon[i+10] = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
			vector3df(-2745, 475, 5970 + 40 * i), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
	}
	ftoneneon[0] = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
		vector3df(-2645, 300, 6520), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
	fttenneon[0] = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
		vector3df(-2645, 300, 6770), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
	ftsetneon[0] = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
		vector3df(-2745, 245, 6520), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
	ftaddneon[0] = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
		vector3df(-2745, 245, 6612), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
	ftsubneon[0] = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
		vector3df(-2745, 245, 6640), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
	ftringneon[0] = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
		vector3df(-2645, 245, 6737), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
	ftoneneon[1] = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
		vector3df(2922, 300, 11230), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
	fttenneon[1] = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
		vector3df(2922, 300, 10980), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
	ftsetneon[1] = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
		vector3df(3022, 245, 11230), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
	ftaddneon[1] = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
		vector3df(3022, 245, 11138), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
	ftsubneon[1] = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
		vector3df(3022, 245, 11110), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
	ftringneon[1] = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
		vector3df(2922, 245, 11010), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
	ftoneneon[2] = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
		vector3df(2922, 300, 10010), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
	fttenneon[2] = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
		vector3df(2922, 300, 9760), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
	ftsetneon[2] = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
		vector3df(3022, 245, 10010), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
	ftaddneon[2] = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
		vector3df(3022, 245, 9918), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
	ftsubneon[2] = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
		vector3df(3022, 245, 9890), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
	ftringneon[2] = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
		vector3df(2922, 245, 9790), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
	for(i = 0; i < 10; i++) {
		consneon[i] = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 160, 0), SColor(128, 255, 150, 0),
			vector3df(3095, 660, 7570 - 49 * i), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
		consneon[i+10] = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 160, 0), SColor(128, 255, 150, 0),
			vector3df(3095, 204, 7570 - 49 * i), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
	}
	mulsneon = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
		vector3df(-910, 225, 14090), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
	mulr1neon = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
		vector3df(-1400, 230, 14190), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
	mulr3neon = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
		vector3df(-180, 230, 14190), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
	dsqplneon = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
		vector3df(-2720, 190, 8970), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
	for(i = 0; i < 10; i++) {
		dsqstat[i] = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
			vector3df(-2745, 370, 9080 + i * 21), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
		dsqstat[i+10] = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
			vector3df(-2820, 190, 9010 + i * 45), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
		dsqstat[i+20] = smgr->addVolumeLightSceneNode(0, -1, 32, 32, SColor(128, 255, 150, 0), SColor(128, 255, 150, 0),
			vector3df(-2820, 165, 9010 + i * 45), vector3df(0, 0, 0), vector3df(8.0, 8.0, 8.0));
	}

	std::cout << "ready\n" << std::flush;

	while(device->run()) {
		driver->beginScene(true, true, SColor(255, 105, 110, 130));
		smgr->setActiveCamera(camera);
		driver->setViewPort(rect<s32>(20, -200, 1660, 1050));
		smgr->drawAll();
		guienv->drawAll();
		driver->endScene();
		device->yield();
	}
	device->drop();
}
