// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2018 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package disks

import (
	"github.com/snapcore/snapd/osutil"
)

// DiskFromDeviceName is not implemented on darwin
func DiskFromDeviceName(deviceName string) (Disk, error) {
	return nil, osutil.ErrDarwin
}

var diskFromDeviceName = func(mountpoint string) (Disk, error) {
	return nil, osutil.ErrDarwin
}

// DiskFromMountPoint is not implemented on darwin
func DiskFromMountPoint(mountpoint string, opts *Options) (Disk, error) {
	return nil, osutil.ErrDarwin
}

var diskFromMountPoint = func(mountpoint string, opts *Options) (Disk, error) {
	return nil, osutil.ErrDarwin
}

var rootMountPointsForPartition = func(p Partition) ([]string, error) {
	return nil, osutil.ErrDarwin
}
