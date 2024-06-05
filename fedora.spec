Name:           calc
Version:        1.0
Release:        1%{?dist}
Summary:        Simple calc app

License:        GPLv3
URL:            https://github.com/kk7453603/Calc_BSBO_01_21
Source0:        %{name}-%{version}.tar.gz

BuildRequires:  systemd-rpm-macros

%description

%global debug_package %{nil}

%prep
%autosetup


%build
export CGO_CFLAGS="%{optflags}"
export CGO_LDFLAGS="-Wl,--build-id"
go build -v -o "./calc" "./cmd/main.go"

%install
mkdir -p $RPM_BUILD_ROOT/opt/%{name}
install calc $RPM_BUILD_ROOT/opt/%{name}

%files
%defattr(-,root,root)
/opt/%{name}

%clean
rm -rf $RPM_BUILD_ROOT

%changelog
* Sat Mar 16 2024 kk7453603 <kk7453603@gmail.com>
- Initial package release
